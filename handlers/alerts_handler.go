package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tailwarden/komiser/models"
)

func (handler *ApiHandler) IsSlackEnabledHandler(c *gin.Context) {
	output := struct {
		Enabled bool `json:"enabled"`
	}{
		Enabled: false,
	}
	if len(handler.cfg.Slack.Webhook) > 0 {
		output.Enabled = true
	}

	c.JSON(http.StatusOK, output)
}

func (handler *ApiHandler) NewAlertHandler(c *gin.Context) {
	var alert models.Alert

	err := json.NewDecoder(c.Request.Body).Decode(&alert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := handler.db.NewInsert().Model(&alert).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	alertId, _ := result.LastInsertId()
	alert.Id = alertId

	if handler.telemetry {
		handler.analytics.TrackEvent("creating_alert", map[string]interface{}{
			"type":        alert.Type,
			"destination": "Slack",
		})
	}

	c.JSON(http.StatusCreated, alert)
}

func (handler *ApiHandler) UpdateAlertHandler(c *gin.Context) {
	alertId := c.Param("id")

	var alert models.Alert
	err := json.NewDecoder(c.Request.Body).Decode(&alert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	_, err = handler.db.NewUpdate().Model(&alert).Column("name", "type", "budget", "usage").Where("id = ?", alertId).Exec(handler.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, alert)
}

func (handler *ApiHandler) DeleteAlertHandler(c *gin.Context) {
	alertId := c.Param("id")

	alert := new(models.Alert)
	_, err := handler.db.NewDelete().Model(alert).Where("id = ?", alertId).Exec(handler.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "alert has been deleted"})
}
