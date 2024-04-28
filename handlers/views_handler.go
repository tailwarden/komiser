package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/repository"
)

func (handler *ApiHandler) NewViewHandler(c *gin.Context) {
	var view models.View

	err := json.NewDecoder(c.Request.Body).Decode(&view)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := repository.HandleQuery(handler.ctx, handler.db, "INSERT", &view, [][3]string{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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

	_, err := repository.HandleQuery(handler.ctx, handler.db, "LIST", &views, [][3]string{})
	if err != nil {
		logrus.WithError(err).Error("scan failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
		return
	}

	c.JSON(http.StatusOK, views)
}

func (handler *ApiHandler) UpdateViewHandler(c *gin.Context) {
	viewId := c.Param("id")

	var view models.View
	err := json.NewDecoder(c.Request.Body).Decode(&view)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = repository.HandleQuery(handler.ctx, handler.db, "UPDATE_VIEW", &view, [][3]string{{"id", "=", fmt.Sprint(viewId)}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, view)
}

func (handler *ApiHandler) DeleteViewHandler(c *gin.Context) {
	viewId := c.Param("id")

	view := new(models.View)
	_, err := repository.HandleQuery(handler.ctx, handler.db, "DELETE", view, [][3]string{{"id", "=", fmt.Sprint(viewId)}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "view has been deleted"})
}

func (handler *ApiHandler) HideResourcesFromViewHandler(c *gin.Context) {
	viewId := c.Param("id")

	var view models.View
	err := json.NewDecoder(c.Request.Body).Decode(&view)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = repository.HandleQuery(handler.ctx, handler.db, "UPDATE_VIEW_EXCLUDE", &view, [][3]string{{"id", "=", fmt.Sprint(viewId)}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "resource has been hidden"})
}

func (handler *ApiHandler) UnhideResourcesFromViewHandler(c *gin.Context) {
	viewId := c.Param("id")

	var view models.View
	err := json.NewDecoder(c.Request.Body).Decode(&view)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = repository.HandleQuery(handler.ctx, handler.db, "UPDATE_VIEW_EXCLUDE", &view, [][3]string{{"id", "=", fmt.Sprint(viewId)}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "resource has been revealed"})
}

func (handler *ApiHandler) ListHiddenResourcesHandler(c *gin.Context) {
	viewId := c.Param("id")

	view := new(models.View)
	_, err := repository.HandleQuery(handler.ctx, handler.db, "LIST", &view, [][3]string{{"id", "=", fmt.Sprint(viewId)}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resources := make([]models.Resource, len(view.Exclude))

	if len(view.Exclude) > 0 {
		s, _ := json.Marshal(view.Exclude)

		_, err = repository.HandleQuery(handler.ctx, handler.db, "LIST", &resources, [][3]string{{"id", "IN", strings.Trim(string(s), "[]")}})
		if err != nil {
			logrus.WithError(err).Error("scan failed")
		}

	}

	c.JSON(http.StatusOK, resources)
}

func (handler *ApiHandler) ListViewAlertsHandler(c *gin.Context) {
	viewId := c.Param("id")

	alerts := make([]models.Alert, 0)

	_, err := repository.HandleQuery(handler.ctx, handler.db, "LIST", &alerts, [][3]string{{"view_id", "=", fmt.Sprint(viewId)}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alerts)
}