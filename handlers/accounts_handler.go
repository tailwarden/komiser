package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/utils"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
)

var unsavedAccounts []models.Account

func (handler *ApiHandler) IsOnboardedHandler(c *gin.Context) {
	output := struct {
		Onboarded bool   `json:"onboarded"`
		Status    string `json:"status"`
	}{
		Onboarded: false,
		Status:    "COMPLETE",
	}

	if handler.db == nil {
		output.Status = "PENDING_DATABASE"
		c.JSON(http.StatusOK, output)
		return
	}

	accounts, err := handler.ctrl.ListAccounts(c)
	if err != nil {
		log.WithError(err).Error("scan failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
	}

	if len(accounts) > 0 {
		output.Onboarded = true
	} else {
		output.Status = "PENDING_ACCOUNTS"
	}
	c.JSON(http.StatusOK, output)
}

func (handler *ApiHandler) ListCloudAccountsHandler(c *gin.Context) {
	if handler.db == nil {
		c.JSON(http.StatusOK, unsavedAccounts)
		return
	}

	accounts, err := handler.ctrl.ListAccounts(c)
	if err != nil {
		log.WithError(err).Error("scan failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
		return
	}

	for i, account := range accounts {
		output, err := handler.ctrl.CountResources(c, account.Provider, account.Name)
		if err != nil {
			log.WithError(err).Error("scan failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
			return
		}
		accounts[i].Resources = output.Total

		if account.Status == "" {
			accounts[i].Status = "CONNECTED"
		}
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

	if handler.db == nil {
		if len(unsavedAccounts) == 0 {
			unsavedAccounts = make([]models.Account, 0)
		}

		unsavedAccounts = append(unsavedAccounts, account)
	}

	accountId, err := handler.ctrl.InsertAccount(c, account)
	if err != nil {
		log.WithError(err).Error("insert failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	account.Id = accountId

	err = populateConfigFromAccount(account, &handler.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = updateConfig(handler.configPath, &handler.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cron := gocron.NewScheduler(time.UTC)
	_, err = cron.Every(1).Hours().Do(func() {
		log.Info("Fetching resources workflow has started")

		fetchResourcesForAccount(c, account, handler.db, []string{})
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cron.StartAsync()

	if handler.telemetry {
		handler.analytics.TrackEvent("creating_alert", map[string]interface{}{
			"type":     len(account.Credentials),
			"provider": account.Provider,
		})
	}

	c.JSON(http.StatusCreated, account)
}

func (handler *ApiHandler) ReScanAccount(c *gin.Context) {
	accountId := c.Param("id")

	account := new(models.Account)
	account.Status = "SCANNING"
	rows, err := handler.ctrl.RescanAccount(c, account, accountId)
	if err != nil {
		log.Error("Couldn't set status", err)
		return
	}
	if rows > 0 {
		go fetchResourcesForAccount(handler.ctx, *account, handler.db, []string{})
	}

	c.JSON(http.StatusOK, "Rescan Triggered")
}

func (handler *ApiHandler) DeleteCloudAccountHandler(c *gin.Context) {
	accountId := c.Param("id")

	res, err := handler.ctrl.GetAccountById(c, accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = handler.ctrl.DeleteAccount(c, accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = deleteConfigAccounts(res, &handler.cfg)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = updateConfig(handler.configPath, &handler.cfg)
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

	err = handler.ctrl.UpdateAccount(c, account, accountId)
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
		sqldb, err := sql.Open(sqliteshim.ShimName, fmt.Sprintf("file:%s?cache=shared", db.FilePath))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		sqldb.SetMaxIdleConns(1000)
		sqldb.SetConnMaxLifetime(0)

		handler.db = bun.NewDB(sqldb, sqlitedialect.New())
		log.Println("Data will be stored in SQLite")

		handler.cfg.SQLite = models.SQLiteConfig{
			File: db.FilePath,
		}
	} else {
		uri := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", db.Username, db.Password, db.Hostname, db.Database)
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(uri)))
		handler.db = bun.NewDB(sqldb, pgdialect.New())

		log.Println("Data will be stored in PostgreSQL")

		handler.cfg.Postgres = models.PostgresConfig{
			URI: uri,
		}
	}

	if len(unsavedAccounts) > 0 {
		if len(handler.accounts) == 0 {
			handler.accounts = unsavedAccounts
		} else {
			handler.accounts = append(handler.accounts, unsavedAccounts...)
		}
		unsavedAccounts = make([]models.Account, 0)
	}

	err = updateConfig(handler.configPath, &handler.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = utils.SetupSchema(handler.db, &handler.cfg, handler.accounts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"message": "database has been configured"})
}
