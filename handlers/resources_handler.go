package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

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
	resources, err := handler.ctrl.ResourceWithFilter(c, *view, []int64{limit, skip}, queryParameter)
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

	view := new(models.View)
		view.Filters = filters

	output, err := handler.ctrl.RelationWithFilter(c, *view, []int64{}, "")
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
