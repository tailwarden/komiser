package internal

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/hashicorp/go-version"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	v1 "github.com/tailwarden/komiser/internal/api/v1"
	"github.com/tailwarden/komiser/internal/config"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/aws"
	azure "github.com/tailwarden/komiser/providers/azure"
	civo "github.com/tailwarden/komiser/providers/civo"
	do "github.com/tailwarden/komiser/providers/digitalocean"
	k8s "github.com/tailwarden/komiser/providers/k8s"
	linode "github.com/tailwarden/komiser/providers/linode"
	oci "github.com/tailwarden/komiser/providers/oci"
	scaleway "github.com/tailwarden/komiser/providers/scaleway"
	"github.com/tailwarden/komiser/providers/tencent"
	"github.com/uptrace/bun"
)

var Version = "Unknown"
var GoVersion = runtime.Version()
var Buildtime = "Unknown"
var Commit = "Unknown"
var Os = runtime.GOOS
var Arch = runtime.GOARCH
var db *bun.DB
var analytics Analytics

func Exec(address string, port int, configPath string, telemetry bool, a Analytics, regions []string, cmd *cobra.Command) error {
	cfg, clients, err := config.Load(configPath)
	if err != nil {
		return err
	}

	analytics = a

	err = setupSchema(cfg)
	if err != nil {
		return err
	}

	cron := gocron.NewScheduler(time.UTC)

	cron.Every(1).Hours().Do(func() {
		log.Info("Fetching resources workflow has started")
		err = fetchResources(context.Background(), clients, regions, telemetry)
		if err != nil {
			log.Fatal(err)
		}
	})

	cron.StartAsync()

	go checkUpgrade()

	err = runServer(address, port, telemetry, *cfg)
	if err != nil {
		return err
	}

	return nil
}

func runServer(address string, port int, telemetry bool, cfg models.Config) error {
	log.Infof("Komiser version: %s, commit: %s, buildt: %s", Version, Commit, Buildtime)

	r := v1.Endpoints(context.Background(), telemetry, db, cfg)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"profile", "X-Requested-With", "Content-Type", "Authorization"},
	})

	loggedRouter := handlers.LoggingHandler(os.Stdout, cors.Handler(r))
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), loggedRouter)
	if err != nil {
		return err
	} else {
		log.Info("Server started on %s:%d", address, port)
	}

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
			models.Filter{
				Field:    "tags",
				Operator: "IS_EMPTY",
				Values:   []string{},
			},
		},
	}

	_, err = db.NewInsert().Model(&untaggedResourcesView).Exec(context.Background())
	if err != nil {
		return err
	}

	expensiveResourcesView := models.View{
		Name: "Expensive resources",
		Filters: []models.Filter{
			models.Filter{
				Field:    "cost",
				Operator: "GREATER_THAN",
				Values:   []string{"0"},
			},
		},
	}

	_, err = db.NewInsert().Model(&expensiveResourcesView).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func fetchResources(ctx context.Context, clients []providers.ProviderClient, regions []string, telemetry bool) error {
	for _, client := range clients {
		if client.AWSClient != nil {
			go func(ctx context.Context, client providers.ProviderClient, regions []string) {
				if telemetry {
					analytics.TrackEvent("fetching_resources", map[string]interface{}{
						"provider": "AWS",
					})
				}
				aws.FetchResources(ctx, client, regions, db)
			}(ctx, client, regions)
		} else if client.DigitalOceanClient != nil {
			go func(ctx context.Context, client providers.ProviderClient) {
				if telemetry {
					analytics.TrackEvent("fetching_resources", map[string]interface{}{
						"provider": "DigitalOcean",
					})
				}
				do.FetchResources(ctx, client, db)
			}(ctx, client)
		} else if client.OciClient != nil {
			go func(ctx context.Context, client providers.ProviderClient) {
				if telemetry {
					analytics.TrackEvent("fetching_resources", map[string]interface{}{
						"provider": "OCI",
					})
				}
				oci.FetchResources(ctx, client, db)
			}(ctx, client)
		} else if client.CivoClient != nil {
			go func(ctx context.Context, client providers.ProviderClient) {
				if telemetry {
					analytics.TrackEvent("fetching_resources", map[string]interface{}{
						"provider": "Civo",
					})
				}
				civo.FetchResources(ctx, client, db)
			}(ctx, client)
		} else if client.K8sClient != nil {
			go func(ctx context.Context, client providers.ProviderClient) {
				if telemetry {
					analytics.TrackEvent("fetching_resources", map[string]interface{}{
						"provider": "Kubernetes",
					})
				}
				k8s.FetchResources(ctx, client, db)
			}(ctx, client)
		} else if client.LinodeClient != nil {
			go func(ctx context.Context, client providers.ProviderClient) {
				if telemetry {
					analytics.TrackEvent("fetching_resources", map[string]interface{}{
						"provider": "Linode",
					})
				}
				linode.FetchResources(ctx, client, db)
			}(ctx, client)
		} else if client.TencentClient != nil {
			go func(ctx context.Context, client providers.ProviderClient) {
				if telemetry {
					analytics.TrackEvent("fetching_resources", map[string]interface{}{
						"provider": "Tencent",
					})
				}
				tencent.FetchResources(ctx, client, db)
			}(ctx, client)
		} else if client.AzureClient != nil {
			go func(ctx context.Context, client providers.ProviderClient) {
				if telemetry {
					analytics.TrackEvent("fetching_resources", map[string]interface{}{
						"provider": "Azure",
					})
				}
				azure.FetchResources(ctx, client, db)
			}(ctx, client)
		} else if client.ScalewayClient != nil {
			go func(ctx context.Context, client providers.ProviderClient) {
				if telemetry {
					analytics.TrackEvent("fetching_resources", map[string]interface{}{
						"provider": "Scaleway",
					})
				}
				scaleway.FetchResources(ctx, client, db)
			}(ctx, client)
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
