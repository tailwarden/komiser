package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
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

	accounts := make([]models.Account, 0)
	err := handler.db.NewRaw("SELECT * FROM accounts").Scan(handler.ctx, &accounts)
	if err != nil {
		log.WithError(err).Error("scan failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
		return
	}

	if len(accounts) > 0 {
		output.Onboarded = true
	} else {
		output.Status = "PENDING_ACCOUNTS"
	}

	c.JSON(http.StatusOK, output)
}

func (handler *ApiHandler) ListCloudAccountsHandler(c *gin.Context) {
	accounts := make([]models.Account, 0)

	if handler.db == nil {
		c.JSON(http.StatusOK, unsavedAccounts)
		return
	}

	err := handler.db.NewRaw("SELECT * FROM accounts").Scan(handler.ctx, &accounts)
	if err != nil {
		log.WithError(err).Error("scan failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
		return
	}

	for i, account := range accounts {
		output := struct {
			Total int `bun:"total" json:"total"`
		}{}
		err = handler.db.NewRaw(fmt.Sprintf("SELECT COUNT(*) as total FROM resources WHERE provider='%s' AND account='%s'", account.Provider, account.Name)).Scan(handler.ctx, &output)
		if err != nil {
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
	} else {
		result, err := handler.db.NewInsert().Model(&account).Exec(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		accountId, _ := result.LastInsertId()
		account.Id = accountId
		go fetchResourcesForAccount(c, account, handler.db, []string{})
	}

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

	config := models.Config{}

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

		config.SQLite = models.SQLiteConfig{
			File: db.FilePath,
		}
	} else {
		uri := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", db.Username, db.Password, db.Hostname, db.Database)
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(uri)))
		handler.db = bun.NewDB(sqldb, pgdialect.New())

		log.Println("Data will be stored in PostgreSQL")

		config.Postgres = models.PostgresConfig{
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

		f, err := os.Create("config.toml")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := toml.NewEncoder(f).Encode(config); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := f.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	err = utils.SetupSchema(handler.db, &handler.cfg, handler.accounts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"message": "database has been configured"})
}
