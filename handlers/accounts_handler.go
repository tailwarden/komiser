package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
)

func (handler *ApiHandler) IsOnboardedHandler(c *gin.Context) {
	output := struct {
		Onboarded bool `json:"onboarded"`
	}{
		Onboarded: false,
	}

	accounts := make([]models.Account, 0)
	err := handler.db.NewRaw("SELECT * FROM accounts").Scan(handler.ctx, &accounts)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
		return
	}

	if len(accounts) > 0 {
		output.Onboarded = true
	}

	c.JSON(http.StatusOK, output)
}

func (handler *ApiHandler) ListCloudAccountsHandler(c *gin.Context) {
	accounts := make([]models.Account, 0)
	err := handler.db.NewRaw("SELECT * FROM accounts").Scan(handler.ctx, &accounts)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func (handler *ApiHandler) NewCloudAccountHandler(c *gin.Context) {
	var account models.Account

	err := json.NewDecoder(c.Request.Body).Decode(&account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.db.NewInsert().Model(&account).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accountId, _ := result.LastInsertId()
	account.Id = accountId

	if handler.telemetry {
		handler.analytics.TrackEvent("creating_alert", map[string]interface{}{
			"type":     len(account.Credentials),
			"provider": account.Provider,
		})
	}

	c.JSON(http.StatusCreated, account)
}
