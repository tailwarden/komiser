package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func (handler *ApiHandler) IsOnboardedHandler(c *gin.Context) {
	output := struct {
		Onboarded bool `json:"onboarded"`
	}{
		Onboarded: false,
	}

	if handler.db != nil {
		c.JSON(http.StatusOK, output)
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

func (handler *ApiHandler) DeleteCloudAccountHandler(c *gin.Context) {
	accountId := c.Param("id")

	account := new(models.Account)
	_, err := handler.db.NewDelete().Model(account).Where("id = ?", accountId).Exec(handler.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "account has been deleted"})
}

func (handler *ApiHandler) UpdateCloudAccountHandler(c *gin.Context) {
	accountId := c.Param("id")

	var account models.Account
	err := json.NewDecoder(c.Request.Body).Decode(&account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = handler.db.NewUpdate().Model(&account).Column("name", "provider", "credentials").Where("id = ?", accountId).Exec(handler.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (handler *ApiHandler) ConfigureDatabaseHandler(c *gin.Context) {
	var db models.DatabaseConfig
	err := json.NewDecoder(c.Request.Body).Decode(&db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if db.Type == "SQLITE" {
		sqldb, err := sql.Open(sqliteshim.ShimName, fmt.Sprintf("file:%s?cache=shared", db.File))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		sqldb.SetMaxIdleConns(1000)
		sqldb.SetConnMaxLifetime(0)

		handler.db = bun.NewDB(sqldb, sqlitedialect.New())
		log.Println("Data will be stored in SQLite")
	} else {
		uri := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", db.Username, db.Password, db.Hostname, db.Database)
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(uri)))
		handler.db = bun.NewDB(sqldb, pgdialect.New())

		log.Println("Data will be stored in PostgreSQL")
	}

	c.JSON(http.StatusOK, map[string]string{"message": "database has been configured"})
}
