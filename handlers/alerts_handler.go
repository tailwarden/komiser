package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

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
		return
	}

	result, err := handler.db.NewInsert().Model(&alert).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
		return
	}

	_, err = handler.db.NewUpdate().Model(&alert).Column("name", "type", "budget", "usage", "endpoint", "secret").Where("id = ?", alertId).Exec(handler.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alert)
}

func (handler *ApiHandler) DeleteAlertHandler(c *gin.Context) {
	alertId := c.Param("id")

	alert := new(models.Alert)
	_, err := handler.db.NewDelete().Model(alert).Where("id = ?", alertId).Exec(handler.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "alert has been deleted"})
}

func (handler *ApiHandler) TestEndpointHandler(c *gin.Context) {
	var endpoint models.Endpoint

	err := json.NewDecoder(c.Request.Body).Decode(&endpoint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	var payloadJSON []byte
	payload := models.CustomWebhookPayload{
		Komiser:   "komiser version that will send the webhook",
		View:      "testing the connection",
		Message:   "test alert",
		Data:      0,
		Timestamp: time.Now().Unix(),
	}

	payloadJSON, err = json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	req, err := http.NewRequest("POST", endpoint.Url, bytes.NewBuffer(payloadJSON))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errMessage := "Custom Webhook with endpoint " + endpoint.Url + " returned back a status code of " + string(rune(resp.StatusCode)) + " . Expected Status Code: 200"
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": errMessage})
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Pinged server successfully"})

}
