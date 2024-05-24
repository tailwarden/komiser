package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
)

func (handler *ApiHandler) NewViewHandler(c *gin.Context) {
	var view models.View

	err := json.NewDecoder(c.Request.Body).Decode(&view)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	viewId, err := handler.ctrl.InsertView(c, view)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	view.Id = viewId

	if handler.telemetry {
		handler.analytics.TrackEvent("creating_view", nil)
	}

	c.JSON(http.StatusCreated, view)
}

func (handler *ApiHandler) ListViewsHandler(c *gin.Context) {
	views, err := handler.ctrl.ListViews(c)
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

	err = handler.ctrl.UpdateView(c, view, viewId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, view)
}

func (handler *ApiHandler) DeleteViewHandler(c *gin.Context) {
	viewId := c.Param("id")

	err := handler.ctrl.DeleteView(c, viewId)
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

	err = handler.ctrl.UpdateViewExclude(c, view, viewId)
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
	err = handler.ctrl.UpdateViewExclude(c, view, viewId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "resource has been revealed"})
}

func (handler *ApiHandler) ListHiddenResourcesHandler(c *gin.Context) {
	viewId := c.Param("id")

	view, err := handler.ctrl.GetView(c, viewId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resources []models.Resource
	if len(view.Exclude) > 0 {
		s, _ := json.Marshal(view.Exclude)

		resources, err = handler.ctrl.GetResources(c, string(s))
		if err != nil {
			logrus.WithError(err).Error("scan failed")
		}

	}

	c.JSON(http.StatusOK, resources)
}

func (handler *ApiHandler) ListViewAlertsHandler(c *gin.Context) {
	viewId := c.Param("id")

	alerts, err := handler.ctrl.ListViewAlerts(c, viewId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alerts)
}
