package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/utils"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect"
)

type ApiHandler struct {
	db        *bun.DB
	ctx       context.Context
	telemetry bool
	cfg       models.Config
	analytics utils.Analytics
}

func NewApiHandler(ctx context.Context, telemetry bool, analytics utils.Analytics, db *bun.DB, cfg models.Config) *ApiHandler {
	handler := ApiHandler{
		db:        db,
		ctx:       ctx,
		telemetry: telemetry,
		cfg:       cfg,
		analytics: analytics,
	}
	return &handler
}

func (handler *ApiHandler) FilterResourcesHandler(c *gin.Context) {
	var filters []Filter

	limitRaw := c.Query("limit")
	skipRaw := c.Query("skip")
	query := c.Query("query")
	viewId := c.Query("view")

	view := new(View)
	if viewId != "" {
		err := handler.db.NewSelect().Model(view).Where("id = ?", viewId).Scan(handler.ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	var limit int64
	var skip int64
	l, err := strconv.ParseInt(limitRaw, 10, 64)
	if err != nil {
		limit = 0
	} else {
		limit = l
	}

	s, err := strconv.ParseInt(skipRaw, 10, 64)
	if err != nil {
		skip = 0
	} else {
		skip = s
	}

	err = json.NewDecoder(c.Request.Body).Decode(&filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost = %f)", cost))
			case "BETWEEN":
				min, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
				}
				max, err := strconv.ParseFloat(filter.Values[1], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost >= %f AND cost <= %f)", min, max))
			case "GREATER_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost > %f)", cost))
			case "LESS_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost < %f)", cost))
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "field is invalid or not supported"})
			return
		}
	}

	if len(query) > 0 {
		clause := fmt.Sprintf("(name ilike '%s' OR region ilike '%s' OR service ilike '%s' OR provider ilike '%s' OR account ilike '%s' OR tags @> '[{\"value\":\"%s\"}]' or tags @> '[{\"key\":\"%s\"}]')", query, query, query, query, query, query, query)
		whereQueries = append(whereQueries, clause)
	}

	whereClause := strings.Join(whereQueries, " AND ")

	resources := make([]Resource, 0)

	if len(filters) == 0 {
		if len(query) > 0 {
			whereClause := fmt.Sprintf("(name ilike '%s' OR region ilike '%s' OR service ilike '%s' OR provider ilike '%s' OR account ilike '%s' OR tags @> '[{\"value\":\"%s\"}]' or tags @> '[{\"key\":\"%s\"}]')", query, query, query, query, query, query, query)
			err = handler.db.NewRaw(fmt.Sprintf("SELECT * FROM resources WHERE %s ORDER BY id LIMIT %d OFFSET %d", whereClause, limit, skip)).Scan(handler.ctx, &resources)
			if err != nil {
				logrus.WithError(err).Error("scan failed")
			}
		} else {
			err = handler.db.NewRaw(fmt.Sprintf("SELECT * FROM resources ORDER BY id LIMIT %d OFFSET %d", limit, skip)).Scan(handler.ctx, &resources)
			if err != nil {
				logrus.WithError(err).Error("scan failed")
			}
		}
		c.JSON(http.StatusOK, resources)
		return
	}

	if filterWithTags {
		query := fmt.Sprintf("SELECT id, resource_id, provider, account, service, region, name, created_at, fetched_at,cost, metadata, tags,link FROM resources CROSS JOIN jsonb_array_elements(tags) AS res WHERE %s ORDER BY id LIMIT %d OFFSET %d", whereClause, limit, skip)
		if len(view.Exclude) > 0 {
			s, _ := json.Marshal(view.Exclude)
			query = fmt.Sprintf("SELECT id, resource_id, provider, account, service, region, name, created_at, fetched_at,cost, metadata, tags,link FROM resources CROSS JOIN jsonb_array_elements(tags) AS res WHERE %s AND id NOT IN (%s) ORDER BY id LIMIT %d OFFSET %d", whereClause, strings.Trim(string(s), "[]"), limit, skip)
		}
		if handler.db.Dialect().Name() == dialect.SQLite {
			query = fmt.Sprintf("SELECT resources.id, resources.resource_id, resources.provider, resources.account, resources.service, resources.region, resources.name, resources.created_at, resources.fetched_at, resources.cost, resources.metadata, resources.tags, resources.link FROM resources CROSS JOIN json_each(tags) WHERE type='object' AND %s ORDER BY resources.id LIMIT %d OFFSET %d", whereClause, limit, skip)
			if len(view.Exclude) > 0 {
				s, _ := json.Marshal(view.Exclude)
				query = fmt.Sprintf("SELECT resources.id, resources.resource_id, resources.provider, resources.account, resources.service, resources.region, resources.name, resources.created_at, resources.fetched_at, resources.cost, resources.metadata, resources.tags, resources.link FROM resources CROSS JOIN json_each(tags) WHERE resources.id NOT IN (%s) AND type='object' AND %s ORDER BY resources.id LIMIT %d OFFSET %d", strings.Trim(string(s), "[]"), whereClause, limit, skip)
			}
		}

		err = handler.db.NewRaw(query).Scan(handler.ctx, &resources)
		if err != nil {
			logrus.WithError(err).Error("scan failed")
		}
	} else {
		query := fmt.Sprintf("SELECT * FROM resources WHERE %s ORDER BY id LIMIT %d OFFSET %d", whereClause, limit, skip)
		if len(view.Exclude) > 0 {
			s, _ := json.Marshal(view.Exclude)
			query = fmt.Sprintf("SELECT * FROM resources WHERE %s AND id NOT IN (%s) ORDER BY id LIMIT %d OFFSET %d", whereClause, strings.Trim(string(s), "[]"), limit, skip)
		}

		err = handler.db.NewRaw(query).Scan(handler.ctx, &resources)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, resources)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		logrus.WithError(err).Error("write failed")
	}
}
