package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	. "github.com/tailwarden/komiser/models"
	"github.com/uptrace/bun/dialect"
)

func (handler *ApiHandler) StatsHandler(c *gin.Context) {
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

	c.JSON(http.StatusOK, output)
}

func (handler *ApiHandler) FilterStatsHandler(c *gin.Context) {
	var filters []Filter

	err := json.NewDecoder(c.Request.Body).Decode(&filters)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"message": err.Error()})
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
				c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
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
				c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
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
				c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
				return
			}
		} else if filter.Field == "cost" {
			switch filter.Operator {
			case "EQUAL":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost = %f)", cost))
			case "BETWEEN":
				min, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
				}
				max, err := strconv.ParseFloat(filter.Values[1], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost >= %f AND cost <= %f)", min, max))
			case "GREATER_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost > %f)", cost))
			case "LESS_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost < %f)", cost))
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
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

		c.JSON(http.StatusOK, output)
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

		c.JSON(http.StatusOK, output)
	}
}

func (handler *ApiHandler) ListRegionsHandler(c *gin.Context) {
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

	c.JSON(http.StatusOK, regions)
}

func (handler *ApiHandler) ListProvidersHandler(c *gin.Context) {
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

	c.JSON(http.StatusOK, providers)
}

func (handler *ApiHandler) ListServicesHandler(c *gin.Context) {
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

	c.JSON(http.StatusOK, services)
}

func (handler *ApiHandler) ListAccountsHandler(c *gin.Context) {
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

	c.JSON(http.StatusOK, accounts)
}
