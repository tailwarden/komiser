package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	. "github.com/tailwarden/komiser/models"
	"github.com/uptrace/bun/dialect"
)

func (handler *ApiHandler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	regions := struct {
		Count int `bun:"count" json:"total"`
	}{}

	err := handler.db.NewRaw("SELECT COUNT(*) as count FROM (SELECT DISTINCT region FROM resources) AS temp").Scan(handler.ctx, &regions)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	resources := struct {
		Count int `bun:"count" json:"total"`
	}{}

	err = handler.db.NewRaw("SELECT COUNT(*) as count FROM resources").Scan(handler.ctx, &resources)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	cost := struct {
		Sum float64 `bun:"sum" json:"total"`
	}{}

	err = handler.db.NewRaw("SELECT SUM(cost) as sum FROM resources").Scan(handler.ctx, &cost)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	output := struct {
		Resources int     `json:"resources"`
		Regions   int     `json:"regions"`
		Costs     float64 `json:"costs"`
	}{
		Resources: resources.Count,
		Regions:   regions.Count,
		Costs:     cost.Sum,
	}

	if handler.telemetry {
		handler.analytics.TrackEvent("global_stats", map[string]interface{}{
			"costs":     cost.Sum,
			"regions":   regions.Count,
			"resources": resources.Count,
		})
	}

	respondWithJSON(w, 200, output)
}

func (handler *ApiHandler) FilterStatsHandler(w http.ResponseWriter, r *http.Request) {
	var filters []Filter

	err := json.NewDecoder(r.Body).Decode(&filters)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

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
				respondWithError(w, http.StatusBadRequest, "Operation is invalid or not supported")
				return
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
				if handler.db.Dialect().Name() == dialect.SQLite {
					query = fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') IN (%s)))", key, strings.Join(filter.Values, ","))
				}
				whereQueries = append(whereQueries, query)
			case "NOT_CONTAINS":
			case "IS_NOT":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' NOT IN (%s)))", key, strings.Join(filter.Values, ","))
				if handler.db.Dialect().Name() == dialect.SQLite {
					query = fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') NOT IN (%s)))", key, strings.Join(filter.Values, ","))
				}
				whereQueries = append(whereQueries, query)
			case "IS_EMPTY":
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') = ''))", key))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' = ''))", key))
				}
			case "IS_NOT_EMPTY":
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') != ''))", key))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' != ''))", key))
				}
			default:
				respondWithError(w, http.StatusBadRequest, "Operation is invalid or not supported")
				return
			}
		} else if filter.Field == "tags" {
			switch filter.Operator {
			case "IS_EMPTY":
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, "json_array_length(tags) = 0")
				} else {
					whereQueries = append(whereQueries, "jsonb_array_length(tags) = 0")
				}
			case "IS_NOT_EMPTY":
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, "json_array_length(tags) != 0")
				} else {
					whereQueries = append(whereQueries, "jsonb_array_length(tags) != 0")
				}
			default:
				respondWithError(w, http.StatusBadRequest, "Operation is invalid or not supported")
				return
			}
		} else if filter.Field == "cost" {
			switch filter.Operator {
			case "EQUAL":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, "The value should be a number")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost = %f)", cost))
			case "BETWEEN":
				min, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, "The value should be a number")
				}
				max, err := strconv.ParseFloat(filter.Values[1], 64)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, "The value should be a number")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost >= %f AND cost <= %f)", min, max))
			case "GREATER_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, "The value should be a number")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost > %f)", cost))
			case "LESS_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, "The value should be a number")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost < %f)", cost))
			default:
				respondWithError(w, http.StatusBadRequest, "Operation is invalid or not supported")
				return
			}
		} else {
			respondWithError(w, http.StatusBadRequest, "Field is invalid or not supported")
			return
		}
	}

	whereClause := strings.Join(whereQueries, " AND ")

	if filterWithTags {
		regions := struct {
			Count int `bun:"count" json:"total"`
		}{}

		query := fmt.Sprintf("FROM resources CROSS JOIN jsonb_array_elements(tags) AS res WHERE %s", whereClause)
		if handler.db.Dialect().Name() == dialect.SQLite {
			query = fmt.Sprintf("FROM resources CROSS JOIN json_each(tags) WHERE type='object' AND %s", whereClause)
		}

		err = handler.db.NewRaw(fmt.Sprintf("SELECT COUNT(*) as count FROM (SELECT DISTINCT region %s) AS temp", query)).Scan(handler.ctx, &regions)
		if err != nil {
			logrus.WithError(err).Error("scan failed")
		}

		resources := struct {
			Count int `bun:"count" json:"total"`
		}{}

		err = handler.db.NewRaw(fmt.Sprintf("SELECT COUNT(*) as count %s", query)).Scan(handler.ctx, &resources)
		if err != nil {
			logrus.WithError(err).Error("scan failed")
		}

		cost := struct {
			Sum float64 `bun:"sum" json:"total"`
		}{}

		err = handler.db.NewRaw(fmt.Sprintf("SELECT SUM(cost) as sum %s", query)).Scan(handler.ctx, &cost)
		if err != nil {
			logrus.WithError(err).Error("scan failed")
		}

		output := struct {
			Resources int     `json:"resources"`
			Regions   int     `json:"regions"`
			Costs     float64 `json:"costs"`
		}{
			Resources: resources.Count,
			Regions:   regions.Count,
			Costs:     cost.Sum,
		}

		respondWithJSON(w, 200, output)
	} else {
		query := fmt.Sprintf("FROM resources WHERE %s", whereClause)

		regions := struct {
			Count int `bun:"count" json:"total"`
		}{}
		err = handler.db.NewRaw(fmt.Sprintf("SELECT COUNT(*) as count FROM (SELECT DISTINCT region %s) AS temp", query)).Scan(handler.ctx, &regions)
		if err != nil {
			logrus.WithError(err).Error("scan failed")
		}

		resources := struct {
			Count int `bun:"count" json:"total"`
		}{}

		err = handler.db.NewRaw(fmt.Sprintf("SELECT COUNT(*) as count %s", query)).Scan(handler.ctx, &resources)
		if err != nil {
			logrus.WithError(err).Error("scan failed")
		}

		cost := struct {
			Sum float64 `bun:"sum" json:"total"`
		}{}

		err = handler.db.NewRaw(fmt.Sprintf("SELECT SUM(cost) as sum %s", query)).Scan(handler.ctx, &cost)
		if err != nil {
			logrus.WithError(err).Error("scan failed")
		}

		output := struct {
			Resources int     `json:"resources"`
			Regions   int     `json:"regions"`
			Costs     float64 `json:"costs"`
		}{
			Resources: resources.Count,
			Regions:   regions.Count,
			Costs:     cost.Sum,
		}

		respondWithJSON(w, 200, output)
	}
}

func (handler *ApiHandler) ListRegionsHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Region string `bun:"region" json:"region"`
	}

	outputs := make([]Output, 0)

	err := handler.db.NewRaw("SELECT DISTINCT(region) FROM resources").Scan(handler.ctx, &outputs)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	regions := make([]string, 0)

	for _, o := range outputs {
		regions = append(regions, o.Region)
	}

	respondWithJSON(w, 200, regions)
}

func (handler *ApiHandler) ListProvidersHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Provider string `bun:"provider" json:"provider"`
	}

	outputs := make([]Output, 0)

	err := handler.db.NewRaw("SELECT DISTINCT(provider) FROM resources").Scan(handler.ctx, &outputs)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	providers := make([]string, 0)

	for _, o := range outputs {
		providers = append(providers, o.Provider)
	}

	respondWithJSON(w, 200, providers)
}

func (handler *ApiHandler) ListServicesHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Service string `bun:"service" json:"service"`
	}

	outputs := make([]Output, 0)

	err := handler.db.NewRaw("SELECT DISTINCT(service) FROM resources").Scan(handler.ctx, &outputs)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	services := make([]string, 0)

	for _, o := range outputs {
		services = append(services, o.Service)
	}

	respondWithJSON(w, 200, services)
}

func (handler *ApiHandler) ListAccountsHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Account string `bun:"account" json:"account"`
	}

	outputs := make([]Output, 0)

	err := handler.db.NewRaw("SELECT DISTINCT(account) FROM resources").Scan(handler.ctx, &outputs)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	accounts := make([]string, 0)

	for _, o := range outputs {
		accounts = append(accounts, o.Account)
	}

	respondWithJSON(w, 200, accounts)
}
