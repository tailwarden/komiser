package utils

import (
	"context"
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/migrations"
	"github.com/tailwarden/komiser/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/migrate"
)

func SetupDBConnection(c *models.Config) (*bun.DB, error) {
	var sqldb *sql.DB
	var err error
	var db *bun.DB

	if len(c.SQLite.File) == 0 && len(c.Postgres.URI) == 0 {
		log.Println("Database wasn't configured yet")
		return nil, nil
	}

	if len(c.SQLite.File) > 0 {
		sqldb, err = sql.Open(sqliteshim.ShimName, fmt.Sprintf("file:%s?cache=shared", c.SQLite.File))
		if err != nil {
			return nil, err
		}
		sqldb.SetMaxIdleConns(1000)
		sqldb.SetConnMaxLifetime(0)

		db = bun.NewDB(sqldb, sqlitedialect.New())

		log.Println("Data will be stored in SQLite")
	} else {
		sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(c.Postgres.URI)))
		db = bun.NewDB(sqldb, pgdialect.New())
		log.Println("Data will be stored in PostgreSQL")
	}

	return db, nil
}

func SetupSchema(db *bun.DB, c *models.Config, accounts []models.Account) error {
	_, err := db.NewCreateTable().Model((*models.Resource)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	_, err = db.NewCreateTable().Model((*models.View)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	_, err = db.NewCreateTable().Model((*models.Alert)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	_, err = db.NewCreateTable().Model((*models.Account)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	for _, account := range accounts {
		account.Status = "CONNECTED"
		_, err = db.NewInsert().Model(&account).Exec(context.Background())
		if err != nil {
			log.Warnf("%s account cannot be inserted to database\n%v", account.Provider, err)
		}
	}

	// Created pre-defined views
	untaggedResourcesView := models.View{
		Name: "Untagged resources",
		Filters: []models.Filter{
			{
				Field:    "tags",
				Operator: "IS_EMPTY",
				Values:   []string{},
			},
		},
	}

	count, _ := db.NewSelect().Model(&untaggedResourcesView).Where("name = ?", untaggedResourcesView.Name).ScanAndCount(context.Background())
	if count == 0 {
		_, err = db.NewInsert().Model(&untaggedResourcesView).Exec(context.Background())
		if err != nil {
			return err
		}
	}

	expensiveResourcesView := models.View{
		Name: "Expensive resources",
		Filters: []models.Filter{
			{
				Field:    "cost",
				Operator: "GREATER_THAN",
				Values:   []string{"0"},
			},
		},
	}

	count, _ = db.NewSelect().Model(&expensiveResourcesView).Where("name = ?", expensiveResourcesView.Name).ScanAndCount(context.Background())
	if count == 0 {
		_, err = db.NewInsert().Model(&expensiveResourcesView).Exec(context.Background())
		if err != nil {
			return err
		}
	}

	err = doMigrations(db, context.Background())
	if err != nil {
		return err
	}

	return nil
}

func doMigrations(db *bun.DB, ctx context.Context) error {
	migrator := migrate.NewMigrator(db, migrations.Migrations)

	if err := migrator.Init(ctx); err != nil {
		return err
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		log.Infof("there are no new migrations to run (database is up to date)\n")
		return nil
	}
	log.Infof("migrated to %s\n", group)
	return nil
}
