package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tailwarden/komiser/controller"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/repository/postgres"
	"github.com/tailwarden/komiser/repository/sqlite"
	"github.com/tailwarden/komiser/utils"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect"
)

type ApiHandler struct {
	db         *bun.DB
	ctrl       *controller.Controller
	ctx        context.Context
	telemetry  bool
	cfg        models.Config
	configPath string
	analytics  utils.Analytics
	accounts   []models.Account
}

func NewApiHandler(ctx context.Context, telemetry bool, analytics utils.Analytics, db *bun.DB, cfg models.Config, configPath string, accounts []models.Account) *ApiHandler {
	var repo controller.Repository
	if db.Dialect().Name() == dialect.SQLite {
		repo = sqlite.NewRepository(db)
	} else {
		repo = postgres.NewRepository(db)
	}

	handler := ApiHandler{
		db:         db,
		ctrl:       controller.New(repo),
		ctx:        ctx,
		telemetry:  telemetry,
		cfg:        cfg,
		configPath: configPath,
		analytics:  analytics,
		accounts:   accounts,
	}
	return &handler
}

func (handler *ApiHandler) FilterResourcesHandler(c *gin.Context) {
	var filters []models.Filter
	resources := make([]models.Resource, 0)

	limitRaw := c.Query("limit")
	skipRaw := c.Query("skip")
	query := c.Query("query")
	viewId := c.Query("view")
	queryParameter := query

	view := new(models.View)
	if viewId != "" {
		err := handler.db.NewSelect().Model(view).Where("id = ?", viewId).Scan(handler.ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
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
	view.Filters = filters
	resources, err = handler.ctrl.ResourceWithFilter(c, *view, []int64{limit, skip}, queryParameter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}		 
	c.JSON(http.StatusOK, resources)
}

func (handler *ApiHandler) RelationStatsHandler(c *gin.Context) {
	var filters []models.Filter

	err := json.NewDecoder(c.Request.Body).Decode(&filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	whereQueries := make([]string, 0)
	for _, filter := range filters {
		if filter.Field == "region" || filter.Field == "service" || filter.Field == "provider" {
			switch filter.Operator {
			case "IS":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("(resources.%s IN (%s))", filter.Field, strings.Join(filter.Values, ","))
				whereQueries = append(whereQueries, query)
			case "IS_NOT":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("(resources.%s NOT IN (%s))", filter.Field, strings.Join(filter.Values, ","))
				whereQueries = append(whereQueries, query)
			case "CONTAINS":
				queries := make([]string, 0)
				specialChar := "%"
				for i := 0; i < len(filter.Values); i++ {
					queries = append(queries, fmt.Sprintf("(resources.%s LIKE '%s%s%s')", filter.Field, specialChar, filter.Values[i], specialChar))
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(%s)", strings.Join(queries, " OR ")))
			case "NOT_CONTAINS":
				queries := make([]string, 0)
				specialChar := "%"
				for i := 0; i < len(filter.Values); i++ {
					queries = append(queries, fmt.Sprintf("(resources.%s NOT LIKE '%s%s%s')", filter.Field, specialChar, filter.Values[i], specialChar))
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(%s)", strings.Join(queries, " AND ")))
			case "IS_EMPTY":
				whereQueries = append(whereQueries, fmt.Sprintf("((coalesce(resources.%s, '') = ''))", filter.Field))
			case "IS_NOT_EMPTY":
				whereQueries = append(whereQueries, fmt.Sprintf("((coalesce(resources.%s, '') != ''))", filter.Field))
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
				return
			}
		} else if filter.Field == "relations" {
			switch filter.Operator {
			case "EQUAL":
				relations, err := strconv.Atoi(filter.Values[0])
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
					return
				}
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, fmt.Sprintf("json_array_length(resources.relations) = %d", relations))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("jsonb_array_length(resources.relations) = %d", relations))
				}
			case "GREATER_THAN":
				relations, err := strconv.Atoi(filter.Values[0])
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
					return
				}
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, fmt.Sprintf("json_array_length(resources.relations) > %d", relations))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("jsonb_array_length(resources.relations) > %d", relations))
				}
			case "LESS_THAN":
				relations, err := strconv.Atoi(filter.Values[0])
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
					return
				}
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, fmt.Sprintf("json_array_length(resources.relations) < %d", relations))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("jsonb_array_length(resources.relations) < %d", relations))
				}
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "field is invalid or not supported"})
			return
		}
	}

	whereClause := strings.Join(whereQueries, " AND ")

	output := make([]models.Resource, 0)

	query := ""
	if len(filters) == 0 {
		query = "SELECT DISTINCT resources.resource_id, resources.provider, resources.name, resources.service, resources.relations FROM resources WHERE (jsonb_array_length(relations) > 0)"
		if handler.db.Dialect().Name() == dialect.SQLite {
			query = "SELECT DISTINCT resources.resource_id, resources.provider, resources.name, resources.service, resources.relations FROM resources WHERE (json_array_length(relations) > 0)"
		}
	} else {
		query = "SELECT DISTINCT resources.resource_id, resources.provider, resources.name, resources.service, resources.relations FROM resources WHERE (jsonb_array_length(relations) > 0) AND " + whereClause
		if handler.db.Dialect().Name() == dialect.SQLite {
			query = "SELECT DISTINCT resources.resource_id, resources.provider, resources.name, resources.service, resources.relations FROM resources WHERE (json_array_length(relations) > 0) AND " + whereClause
		}
	}

	err = handler.db.NewRaw(query).Scan(handler.ctx, &output)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	out := make([]models.OutputRelationResponse, 0)
	for _, ele := range output {
		out = append(out, models.OutputRelationResponse{
			ResourceID: ele.ResourceId,
			Name:       ele.Name,
			Type:       ele.Service,
			Link:       ele.Relations,
			Provider:   ele.Provider,
		})
	}

	c.JSON(http.StatusOK, out)

}

func (handler *ApiHandler) GetResourceByIdHandler(c *gin.Context) {
	resourceId := c.Query("resourceId")

	resource, err := handler.ctrl.GetResource(c, resourceId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		return
	}

	c.JSON(http.StatusOK, resource)
}
