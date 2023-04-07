package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
)

func (handler *ApiHandler) NewViewHandler(c *gin.Context) {
	var view models.View

	err := json.NewDecoder(c.Request.Body).Decode(&view)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := handler.db.NewInsert().Model(&view).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	viewId, _ := result.LastInsertId()
	view.Id = viewId

	if handler.telemetry {
		handler.analytics.TrackEvent("creating_view", nil)
	}

	c.JSON(http.StatusCreated, view)
}

func (handler *ApiHandler) ListViewsHandler(c *gin.Context) {
	views := make([]models.View, 0)

	err := handler.db.NewRaw("SELECT * FROM views").Scan(handler.ctx, &views)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
	}

	c.JSON(http.StatusOK, views)
}

func (handler *ApiHandler) UpdateViewHandler(c *gin.Context) {
	viewId := c.Param("id")

	var view models.View
	err := json.NewDecoder(c.Request.Body).Decode(&view)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	_, err = handler.db.NewUpdate().Model(&view).Column("name", "filters", "exclude").Where("id = ?", viewId).Exec(handler.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, view)
}

func (handler *ApiHandler) DeleteViewHandler(c *gin.Context) {
	viewId := c.Param("id")

	view := new(models.View)
	_, err := handler.db.NewDelete().Model(view).Where("id = ?", viewId).Exec(handler.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "view has been deleted"})
}

func (handler *ApiHandler) HideResourcesFromViewHandler(c *gin.Context) {
	viewId := c.Param("id")

	var view models.View
	err := json.NewDecoder(c.Request.Body).Decode(&view)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	_, err = handler.db.NewUpdate().Model(&view).Column("exclude").Where("id = ?", viewId).Exec(handler.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "resource has been hidden"})
}

func (handler *ApiHandler) UnhideResourcesFromViewHandler(c *gin.Context) {
	viewId := c.Param("id")

	var view models.View
	err := json.NewDecoder(c.Request.Body).Decode(&view)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	_, err = handler.db.NewUpdate().Model(&view).Column("exclude").Where("id = ?", viewId).Exec(handler.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "resource has been revealed"})
}

func (handler *ApiHandler) ListHiddenResourcesHandler(c *gin.Context) {
	viewId := c.Param("id")

	view := new(models.View)
	err := handler.db.NewSelect().Model(view).Where("id = ?", viewId).Scan(handler.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	resources := make([]models.Resource, len(view.Exclude))

	if len(view.Exclude) > 0 {
		s, _ := json.Marshal(view.Exclude)
		err = handler.db.NewRaw(fmt.Sprintf("SELECT * FROM resources WHERE id IN (%s)", strings.Trim(string(s), "[]"))).Scan(handler.ctx, &resources)
		if err != nil {
			logrus.WithError(err).Error("scan failed")
		}

	}

	c.JSON(http.StatusOK, resources)
}

func (handler *ApiHandler) ListViewAlertsHandler(c *gin.Context) {
	viewId := c.Param("id")

	alerts := make([]models.Alert, 0)

	err := handler.db.NewRaw(fmt.Sprintf("SELECT * FROM alerts WHERE view_id = %s", viewId)).Scan(handler.ctx, &alerts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, alerts)
}
