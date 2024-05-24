package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun/dialect"

	"github.com/spf13/cobra"
	"github.com/tailwarden/komiser/internal/config"
	"github.com/tailwarden/komiser/models"

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

func Exec(address string, port int, configPath string, telemetry bool, a utils.Analytics, regions []string, _ *cobra.Command) error {
	analytics = a

	ctx := context.Background()

	cfg, clients, accounts, err := config.Load(configPath, telemetry, analytics)
	if err != nil {
		return err
	}

	db, err = utils.SetupDBConnection(cfg)
	if err != nil {
		return err
	}

	if db != nil {
		err = utils.SetupSchema(db, cfg, accounts)
		if err != nil {
			return err
		}

		scheduleJobs(ctx, cfg, clients, regions, port, telemetry)
	}

	go checkUpgrade()

	err = runServer(address, port, telemetry, *cfg, configPath, accounts)
	if err != nil {
		return err
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
				return models.ViewStat{}, errInvalidOperation
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
				return models.ViewStat{}, errInvalidOperation
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
				return models.ViewStat{}, errInvalidOperation
			}
		} else if filter.Field == "cost" {
			switch filter.Operator {
			case "EQUAL":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					return models.ViewStat{}, errNumberValue
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost = %f)", cost))
			case "BETWEEN":
				min, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					return models.ViewStat{}, errNumberValue
				}
				max, err := strconv.ParseFloat(filter.Values[1], 64)
				if err != nil {
					return models.ViewStat{}, errNumberValue
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost >= %f AND cost <= %f)", min, max))
			case "GREATER_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					return models.ViewStat{}, errNumberValue
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost > %f)", cost))
			case "LESS_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					return models.ViewStat{}, errNumberValue
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost < %f)", cost))
			default:
				return models.ViewStat{}, errInvalidOperation

			}
		} else {
			return models.ViewStat{}, errInvalidField
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
