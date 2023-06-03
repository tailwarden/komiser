package internal

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/hashicorp/go-version"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"github.com/uptrace/bun/dialect"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/migrate"

	"github.com/spf13/cobra"
	v1 "github.com/tailwarden/komiser/internal/api/v1"
	"github.com/tailwarden/komiser/internal/config"
	"github.com/tailwarden/komiser/migrations"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/aws"
	azure "github.com/tailwarden/komiser/providers/azure"
	"github.com/tailwarden/komiser/providers/civo"
	do "github.com/tailwarden/komiser/providers/digitalocean"
	"github.com/tailwarden/komiser/providers/gcp"
	k8s "github.com/tailwarden/komiser/providers/k8s"
	linode "github.com/tailwarden/komiser/providers/linode"
	"github.com/tailwarden/komiser/providers/mongodbatlas"
	"github.com/tailwarden/komiser/providers/oci"
	scaleway "github.com/tailwarden/komiser/providers/scaleway"
	"github.com/tailwarden/komiser/providers/tencent"
	"github.com/tailwarden/komiser/utils"
	"github.com/uptrace/bun"
)

var Version = "Unknown"
var GoVersion = runtime.Version()
var Buildtime = "Unknown"
var Commit = "Unknown"
var Os = runtime.GOOS
var Arch = runtime.GOARCH
var db *bun.DB
var analytics utils.Analytics

func Exec(address string, port int, configPath string, telemetry bool, a utils.Analytics, regions []string, cmd *cobra.Command) error {
	analytics = a

	ctx := context.Background()

	cfg, clients, err := config.Load(configPath, telemetry, analytics)
	if err != nil {
		return err
	}

	err = setupSchema(cfg)
	if err != nil {
		return err
	}

	err = doMigrations(ctx)
	if err != nil {
		return err
	}

	cron := gocron.NewScheduler(time.UTC)

	_, err = cron.Every(1).Hours().Do(func() {
		log.Info("Fetching resources workflow has started")
		err = fetchResources(ctx, clients, regions, telemetry)
		if err != nil {
			log.Fatal(err)
		}
	})

	if err != nil {
		log.WithError(err).Error("setting up cron job failed")
	}

	_, err = cron.Every(1).Hours().Do(func() {
		alertsExist, alerts := checkIfAlertsExist(ctx)

		if alertsExist {
			log.Info("Checking Alerts")
			checkingAlerts(ctx, *cfg, telemetry, port, alerts)
		}
	})

	if err != nil {
		log.WithError(err).Error("setting up cron job failed")
	}

	_, err = cron.Every(1).Friday().At("09:00").Do(func() {
		if len(cfg.Slack.Webhook) > 0 && cfg.Slack.Reporting {
			log.Info("Sending weekly reporting")
			sendTagsCoverageReport(ctx, *cfg)
			sendCostBreakdownReport(ctx, *cfg)
		}
	})

	if err != nil {
		log.WithError(err).Error("setting up cron job failed")
	}

	cron.StartAsync()

	go checkUpgrade()

	err = runServer(address, port, telemetry, *cfg)
	if err != nil {
		return err
	}

	return nil
}

func checkIfAlertsExist(ctx context.Context) (bool, []models.Alert) {
	alerts := make([]models.Alert, 0)

	err := db.NewRaw("SELECT * FROM alerts").Scan(ctx, &alerts)
	if err != nil {
		log.WithError(err).Error("scan failed")
	}
	if len(alerts) > 0 {
		return true, alerts
	}
	return false, alerts
}

func loggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		log.WithFields(log.Fields{
			"method":  reqMethod,
			"uri":     reqUri,
			"status":  statusCode,
			"latency": latencyTime,
			"ip":      clientIP,
		}).Info("HTTP request")

		ctx.Next()
	}
}

func runServer(address string, port int, telemetry bool, cfg models.Config) error {
	log.Infof("Komiser version: %s, commit: %s, buildt: %s", Version, Commit, Buildtime)

	r := v1.Endpoints(context.Background(), telemetry, analytics, db, cfg)

	r.Use(loggingMiddleware())

	if err := r.Run(fmt.Sprintf("%s:%d", address, port)); err != nil {
		return err
	}

	log.Infof("Server started on %s:%d", address, port)

	return nil
}

func setupSchema(c *models.Config) error {
	var sqldb *sql.DB
	var err error

	if len(c.SQLite.File) > 0 {
		sqldb, err = sql.Open(sqliteshim.ShimName, fmt.Sprintf("file:%s?cache=shared", c.SQLite.File))
		if err != nil {
			return err
		}
		sqldb.SetMaxIdleConns(1000)
		sqldb.SetConnMaxLifetime(0)

		db = bun.NewDB(sqldb, sqlitedialect.New())

		log.Println("Data will be stored in SQLite")
	} else {
		sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(c.Postgres.URI)))
		db = bun.NewDB(sqldb, pgdialect.New())

		log.Println("Data will be stored in PostgreSQL")
	}

	_, err = db.NewCreateTable().Model((*models.Resource)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	_, err = db.NewCreateTable().Model((*models.View)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	_, err = db.NewCreateTable().Model((*models.Alert)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	// Created pre-defined views
	untaggedResourcesView := models.View{
		Name: "Untagged resources",
		Filters: []models.Filter{
			{
				Field:    "tags",
				Operator: "IS_EMPTY",
				Values:   []string{},
			},
		},
	}

	count, _ := db.NewSelect().Model(&untaggedResourcesView).Where("name = ?", untaggedResourcesView.Name).ScanAndCount(context.Background())
	if count == 0 {
		_, err = db.NewInsert().Model(&untaggedResourcesView).Exec(context.Background())
		if err != nil {
			return err
		}
	}

	expensiveResourcesView := models.View{
		Name: "Expensive resources",
		Filters: []models.Filter{
			{
				Field:    "cost",
				Operator: "GREATER_THAN",
				Values:   []string{"0"},
			},
		},
	}

	count, _ = db.NewSelect().Model(&expensiveResourcesView).Where("name = ?", expensiveResourcesView.Name).ScanAndCount(context.Background())
	if count == 0 {
		_, err = db.NewInsert().Model(&expensiveResourcesView).Exec(context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}

func doMigrations(ctx context.Context) error {
	migrator := migrate.NewMigrator(db, migrations.Migrations)

	migrator.Init(ctx)

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		log.Infof("there are no new migrations to run (database is up to date)\n")
		return nil
	}
	log.Infof("migrated to %s\n", group)
	return nil
}

func triggerFetchingWorfklow(ctx context.Context, client providers.ProviderClient, provider string, telemetry bool, regions []string) {
	localHub := sentry.CurrentHub().Clone()

	defer func() {
		err := recover()
		log.WithField("err", err).Error(fmt.Sprintf("error fetching %s resources", provider))
		if err != nil {
			localHub.CaptureException(err.(error))
			localHub.Flush(2 * time.Second)
		}
	}()

	localHub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag("provider", provider)
	})

	if telemetry {
		analytics.TrackEvent("fetching_resources", map[string]interface{}{
			"provider": provider,
		})
	}

	switch provider {
	case "AWS":
		aws.FetchResources(ctx, client, regions, db, telemetry, analytics)
	case "DigitalOcean":
		do.FetchResources(ctx, client, db, telemetry, analytics)
	case "OCI":
		oci.FetchResources(ctx, client, db, telemetry, analytics)
	case "Civo":
		civo.FetchResources(ctx, client, db, telemetry, analytics)
	case "Kubernetes":
		k8s.FetchResources(ctx, client, db, telemetry, analytics)
	case "Linode":
		linode.FetchResources(ctx, client, db, telemetry, analytics)
	case "Tencent":
		tencent.FetchResources(ctx, client, db, telemetry, analytics)
	case "Azure":
		azure.FetchResources(ctx, client, db, telemetry, analytics)
	case "Scaleway":
		scaleway.FetchResources(ctx, client, db, telemetry, analytics)
	case "MongoDBAtlas":
		mongodbatlas.FetchResources(ctx, client, db, telemetry, analytics)
	case "GCP":
		gcp.FetchResources(ctx, client, db, telemetry, analytics)
	}
}

func fetchResources(ctx context.Context, clients []providers.ProviderClient, regions []string, telemetry bool) error {
	for _, client := range clients {
		if client.AWSClient != nil {
			go triggerFetchingWorfklow(ctx, client, "AWS", telemetry, regions)
		} else if client.DigitalOceanClient != nil {
			go triggerFetchingWorfklow(ctx, client, "DigitalOcean", telemetry, regions)
		} else if client.OciClient != nil {
			go triggerFetchingWorfklow(ctx, client, "OCI", telemetry, regions)
		} else if client.CivoClient != nil {
			go triggerFetchingWorfklow(ctx, client, "Civo", telemetry, regions)
		} else if client.K8sClient != nil {
			go triggerFetchingWorfklow(ctx, client, "Kubernetes", telemetry, regions)
		} else if client.LinodeClient != nil {
			go triggerFetchingWorfklow(ctx, client, "Linode", telemetry, regions)
		} else if client.TencentClient != nil {
			go triggerFetchingWorfklow(ctx, client, "Tencent", telemetry, regions)
		} else if client.AzureClient != nil {
			go triggerFetchingWorfklow(ctx, client, "Azure", telemetry, regions)
		} else if client.ScalewayClient != nil {
			go triggerFetchingWorfklow(ctx, client, "Scaleway", telemetry, regions)
		} else if client.MongoDBAtlasClient != nil {
			go triggerFetchingWorfklow(ctx, client, "MongoDBAtlas", telemetry, regions)
		} else if client.GCPClient != nil {
			go triggerFetchingWorfklow(ctx, client, "GCP", telemetry, regions)
		}
	}
	return nil
}

func checkUpgrade() {
	url := "https://api.github.com/repos/tailwarden/komiser/releases/latest"
	type GHRelease struct {
		Version string `json:"tag_name"`
	}

	var myClient = &http.Client{Timeout: 5 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		log.Warnf("Failed to check for new version: %s", err)
		return
	}
	defer r.Body.Close()

	target := new(GHRelease)
	err = json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		log.Warnf("Failed to decode new release version: %s", err)
		return
	}

	v1, err := version.NewVersion(Version)
	if err != nil {
		log.Warnf("Failed to parse version: %s", err)
	} else {
		v2, err := version.NewVersion(target.Version)
		if err != nil {
			log.Warnf("Failed to parse version: %s", err)
		} else {
			if v1.LessThan(v2) {
				log.Warnf("Newer Komiser version is available: %s", target.Version)
				log.Warnf("Upgrade instructions: https://github.com/tailwarden/komiser")
			}
		}
	}
}

func hitCustomWebhook(endpoint string, secret string, viewName string, resources int, cost float64, alertType string) {
	var payloadJSON []byte
	var err error
	payload := models.CustomWebhookPayload{
		Komiser:   Version,
		View:      viewName,
		Timestamp: time.Now().Unix(),
	}

	switch alertType {
	case "BUDGET":
		payload.Message = "Cost alert"
		payload.Data = cost
	case "USAGE":
		payload.Message = "Usage alert"
		payload.Data = float64(resources)
	default:
		log.Error("Invalid Alert Type")
		return
	}

	payloadJSON, err = json.Marshal(payload)
	if err != nil {
		log.Error("Couldn't encode JSON payload:", err)
		return
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadJSON))
	if err != nil {
		log.Error("Couldn't create HTTP request for custom webhook endpoint:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if len(secret) > 0 {
		req.Header.Set("Authorization", secret)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Couldn't make HTTP request for custom webhook endpoint:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Error("Custom Webhook with endpoint " + endpoint + " returned back a status code of " + string(rune(resp.StatusCode)) + " . Expected Status Code: 200")
		return
	}
}

func hitSlackWebhook(viewName string, port int, viewId int, resources int, cost float64, webhookUrl string, alertType string) {
	attachment := slack.Attachment{
		Color:         "danger",
		AuthorName:    "Komiser",
		AuthorSubname: "by Tailwarden",
		AuthorLink:    "https://tailwarden.com",
		AuthorIcon:    "https://cdn.komiser.io/images/komiser-logo.jpeg",
		Footer:        "Komiser",
		Actions: []slack.AttachmentAction{
			{
				Name: "open",
				Text: "Open view",
				Type: "button",
				URL:  fmt.Sprintf("http://localhost:%d/inventory?view=%d", port, viewId),
			},
		},
		Fields: []slack.AttachmentField{
			{
				Title: "View",
				Value: viewName,
			},
		},
		FooterIcon: "https://github.com/tailwarden/komiser",
		Ts:         json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}

	if alertType == "BUDGET" {
		attachment.Text = "Cost alert :warning:"
		attachment.Fields = append(attachment.Fields, slack.AttachmentField{
			Title: "Cost",
			Value: fmt.Sprintf("%.2f$", cost),
		})
	} else if alertType == "USAGE" {
		attachment.Text = "Usage alert :warning:"
		attachment.Fields = append(attachment.Fields, slack.AttachmentField{
			Title: "Resources",
			Value: fmt.Sprintf("%d", resources),
		})
	}

	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(webhookUrl, &msg)
	if err != nil {
		log.Warn(err)
	}

}

func checkingAlerts(ctx context.Context, cfg models.Config, telemetry bool, port int, alerts []models.Alert) {
	for _, alert := range alerts {
		var view models.View
		err := db.NewRaw(fmt.Sprintf("SELECT * FROM views WHERE id = %s", alert.ViewId)).Scan(ctx, &view)
		if err != nil {
			log.WithError(err).Error("scan failed")
		}

		stats, err := getViewStats(ctx, view.Filters)
		if err != nil {
			log.Error("Couldn't get stats for view:", view.Name)
		} else {
			if alert.Type == "BUDGET" && alert.Budget <= stats.Costs {
				if telemetry {
					analytics.TrackEvent("sending_alerts", map[string]interface{}{
						"type": "budget",
					})
				}
				if alert.IsSlack {
					log.Info("Sending Slack budget alert for view:", view.Name)
					hitSlackWebhook(view.Name, port, int(view.Id), 0, stats.Costs, cfg.Slack.Webhook, alert.Type)
				} else {
					log.Info("Sending Custom Webhook budget alert for view:", view.Name)
					hitCustomWebhook(alert.Endpoint, alert.Secret, view.Name, 0, stats.Costs, alert.Type)
				}
			}
			if alert.Type == "USAGE" && alert.Usage <= stats.Resources {
				if telemetry {
					analytics.TrackEvent("sending_alerts", map[string]interface{}{
						"type": "usage",
					})
				}
				if alert.IsSlack {
					log.Info("Sending Slack usage alert for view:", view.Name)
					hitSlackWebhook(view.Name, port, int(view.Id), stats.Resources, 0, cfg.Slack.Webhook, alert.Type)
				} else {
					log.Info("Sending Custom Webhook usage alert for view:", view.Name)
					hitCustomWebhook(alert.Endpoint, alert.Secret, view.Name, stats.Resources, 0, alert.Type)
				}
			}
		}
	}
}

func getViewStats(ctx context.Context, filters []models.Filter) (models.ViewStat, error) {
	filterWithTags := false
	whereQueries := make([]string, 0)
	for _, filter := range filters {
		if filter.Field == "name" || filter.Field == "region" || filter.Field == "service" || filter.Field == "provider" || filter.Field == "account" {
			switch filter.Operator {
			case "IS":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("(%s IN (%s))", filter.Field, strings.Join(filter.Values, ","))
				whereQueries = append(whereQueries, query)
			case "IS_NOT":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("(%s NOT IN (%s))", filter.Field, strings.Join(filter.Values, ","))
				whereQueries = append(whereQueries, query)
			case "CONTAINS":
				queries := make([]string, 0)
				specialChar := "%"
				for i := 0; i < len(filter.Values); i++ {
					queries = append(queries, fmt.Sprintf("(%s LIKE '%s%s%s')", filter.Field, specialChar, filter.Values[i], specialChar))
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(%s)", strings.Join(queries, " OR ")))
			case "NOT_CONTAINS":
				queries := make([]string, 0)
				specialChar := "%"
				for i := 0; i < len(filter.Values); i++ {
					queries = append(queries, fmt.Sprintf("(%s NOT LIKE '%s%s%s')", filter.Field, specialChar, filter.Values[i], specialChar))
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(%s)", strings.Join(queries, " AND ")))
			case "IS_EMPTY":
				whereQueries = append(whereQueries, fmt.Sprintf("((coalesce(%s, '') = ''))", filter.Field))
			case "IS_NOT_EMPTY":
				whereQueries = append(whereQueries, fmt.Sprintf("((coalesce(%s, '') != ''))", filter.Field))
			default:
				return models.ViewStat{}, errors.New("Operation is invalid or not supported")
			}
		} else if strings.HasPrefix(filter.Field, "tag:") {
			filterWithTags = true
			key := strings.ReplaceAll(filter.Field, "tag:", "")
			switch filter.Operator {
			case "CONTAINS":
			case "IS":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' IN (%s)))", key, strings.Join(filter.Values, ","))
				if db.Dialect().Name() == dialect.SQLite {
					query = fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') IN (%s)))", key, strings.Join(filter.Values, ","))
				}
				whereQueries = append(whereQueries, query)
			case "NOT_CONTAINS":
			case "IS_NOT":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' NOT IN (%s)))", key, strings.Join(filter.Values, ","))
				if db.Dialect().Name() == dialect.SQLite {
					query = fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') NOT IN (%s)))", key, strings.Join(filter.Values, ","))
				}
				whereQueries = append(whereQueries, query)
			case "IS_EMPTY":
				if db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') = ''))", key))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' = ''))", key))
				}
			case "IS_NOT_EMPTY":
				if db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') != ''))", key))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' != ''))", key))
				}
			default:
				return models.ViewStat{}, errors.New("Operation is invalid or not supported")
			}
		} else if filter.Field == "tags" {
			switch filter.Operator {
			case "IS_EMPTY":
				if db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, "json_array_length(tags) = 0")
				} else {
					whereQueries = append(whereQueries, "jsonb_array_length(tags) = 0")
				}
			case "IS_NOT_EMPTY":
				if db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, "json_array_length(tags) != 0")
				} else {
					whereQueries = append(whereQueries, "jsonb_array_length(tags) != 0")
				}
			default:
				return models.ViewStat{}, errors.New("Operation is invalid or not supported")
			}
		} else if filter.Field == "cost" {
			switch filter.Operator {
			case "EQUAL":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					return models.ViewStat{}, errors.New("The value should be a number")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost = %f)", cost))
			case "BETWEEN":
				min, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					return models.ViewStat{}, errors.New("The value should be a number")
				}
				max, err := strconv.ParseFloat(filter.Values[1], 64)
				if err != nil {
					return models.ViewStat{}, errors.New("The value should be a number")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost >= %f AND cost <= %f)", min, max))
			case "GREATER_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					return models.ViewStat{}, errors.New("The value should be a number")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost > %f)", cost))
			case "LESS_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					return models.ViewStat{}, errors.New("The value should be a number")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost < %f)", cost))
			default:
				return models.ViewStat{}, errors.New("Operation is invalid or not supported")

			}
		} else {
			return models.ViewStat{}, errors.New("Field is invalid or not supported")
		}
	}

	whereClause := strings.Join(whereQueries, " AND ")

	if filterWithTags {
		query := fmt.Sprintf("FROM resources CROSS JOIN jsonb_array_elements(tags) AS res WHERE %s", whereClause)
		if db.Dialect().Name() == dialect.SQLite {
			query = fmt.Sprintf("FROM resources CROSS JOIN json_each(tags) WHERE type='object' AND %s", whereClause)
		}

		resources := struct {
			Count int `bun:"count" json:"total"`
		}{}

		err := db.NewRaw(fmt.Sprintf("SELECT COUNT(*) as count %s", query)).Scan(ctx, &resources)
		if err != nil {
			log.WithError(err).Error("scan failed")
		}

		cost := struct {
			Sum float64 `bun:"sum" json:"total"`
		}{}

		err = db.NewRaw(fmt.Sprintf("SELECT SUM(cost) as sum %s", query)).Scan(ctx, &cost)
		if err != nil {
			log.WithError(err).Error("scan failed")
		}

		output := models.ViewStat{
			Resources: resources.Count,
			Costs:     cost.Sum,
		}

		return output, nil
	} else {
		query := fmt.Sprintf("FROM resources WHERE %s", whereClause)

		resources := struct {
			Count int `bun:"count" json:"total"`
		}{}

		err := db.NewRaw(fmt.Sprintf("SELECT COUNT(*) as count %s", query)).Scan(ctx, &resources)
		if err != nil {
			log.WithError(err).Error("scan failed")
		}

		cost := struct {
			Sum float64 `bun:"sum" json:"total"`
		}{}

		err = db.NewRaw(fmt.Sprintf("SELECT SUM(cost) as sum %s", query)).Scan(ctx, &cost)
		if err != nil {
			log.WithError(err).Error("scan failed")
		}

		output := models.ViewStat{
			Resources: resources.Count,
			Costs:     cost.Sum,
		}

		return output, nil
	}
}

func sendTagsCoverageReport(ctx context.Context, cfg models.Config) {
	tags := make([]struct {
		Total int        `bun:"total"`
		Label models.Tag `bun:"label"`
	}, 0)

	err := db.NewRaw("SELECT count(*) as total, value as label FROM resources CROSS JOIN json_each(tags) GROUP BY value ORDER BY total DESC").Scan(ctx, &tags)
	if err != nil {
		log.WithError(err).Error("scan failed")
	}

	fields := make([]slack.AttachmentField, 0)

	for _, tag := range tags {
		fields = append(fields, slack.AttachmentField{
			Title: fmt.Sprintf("%s:%s", tag.Label.Key, tag.Label.Value),
			Value: fmt.Sprintf("%d", tag.Total),
			Short: true,
		})
	}

	output := struct {
		Total int `bun:"total"`
	}{}

	err = db.NewRaw("SELECT COUNT(*) as total FROM resources where json_array_length(tags) = 0;").Scan(ctx, &output)
	if err != nil {
		log.WithError(err).Error("scan failed")
	}

	currentTime := time.Now()

	attachment := slack.Attachment{
		Color:         "good",
		AuthorName:    "Komiser",
		AuthorSubname: "by Tailwarden",
		AuthorLink:    "https://tailwarden.com",
		AuthorIcon:    "https://cdn.komiser.io/images/komiser-logo.jpeg",
		Text:          fmt.Sprintf("On %s %d: *%d* of your resources are untagged. Below list of most used key/value pairs:", currentTime.Month(), currentTime.Day(), output.Total),
		Footer:        "Komiser",
		Fields:        fields,
		FooterIcon:    "https://github.com/tailwarden/komiser",
		Ts:            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err = slack.PostWebhook(cfg.Slack.Webhook, &msg)
	if err != nil {
		log.Warn(err)
	}
}

func sendCostBreakdownReport(ctx context.Context, cfg models.Config) {
	groups := make([]models.OutputCostByField, 0)
	currentTime := time.Now()

	for _, field := range []string{"service", "provider", "account", "region"} {
		err := db.NewRaw(fmt.Sprintf("SELECT %s as label, SUM(cost) as total FROM resources GROUP BY %s ORDER by total desc;", field, field)).Scan(ctx, &groups)
		if err != nil {
			log.WithError(err).Error("scan failed")
		}

		segments := groups

		if len(groups) > 3 {
			segments = groups[:4]
			if len(groups) > 4 {
				sum := 0.0
				for i := 4; i < len(groups); i++ {
					sum += groups[i].Total
				}

				segments = append(segments, models.OutputCostByField{
					Label: "Others",
					Total: sum,
				})
			}
		}

		fields := make([]slack.AttachmentField, 0)
		for _, segment := range segments {
			fields = append(fields, slack.AttachmentField{
				Title: segment.Label,
				Value: fmt.Sprintf("%.2f", segment.Total),
				Short: true,
			})
		}

		attachment := slack.Attachment{
			Color:         "good",
			AuthorName:    "Komiser",
			AuthorSubname: "by Tailwarden",
			AuthorLink:    "https://tailwarden.com",
			AuthorIcon:    "https://cdn.komiser.io/images/komiser-logo.jpeg",
			Text:          fmt.Sprintf("On %s %d: cost breakdown by cloud %s", currentTime.Month(), currentTime.Day(), field),
			Footer:        "Komiser",
			Fields:        fields,
			FooterIcon:    "https://github.com/tailwarden/komiser",
			Ts:            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
		}
		msg := slack.WebhookMessage{
			Attachments: []slack.Attachment{attachment},
		}

		err = slack.PostWebhook(cfg.Slack.Webhook, &msg)
		if err != nil {
			log.Warn(err)
		}
	}
}
