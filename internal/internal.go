package internal

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/hashicorp/go-version"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"

	"github.com/gorilla/handlers"
	v1 "github.com/mlabouardy/komiser/internal/api/v1"
	"github.com/mlabouardy/komiser/internal/config"
	"github.com/mlabouardy/komiser/models"
	"github.com/mlabouardy/komiser/providers"
	"github.com/mlabouardy/komiser/providers/aws"
	do "github.com/mlabouardy/komiser/providers/digitalocean"
	oci "github.com/mlabouardy/komiser/providers/oci"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
)

var Version = "Unknown"
var GoVersion = runtime.Version()
var Buildtime = "Unknown"
var Commit = "Unknown"
var Os = runtime.GOOS
var Arch = runtime.GOARCH
var db *bun.DB

func Exec(port int, configPath string, noTracking bool, regions []string, cmd *cobra.Command) error {
	cfg, clients, err := config.Load(configPath)
	if err != nil {
		return err
	}

	err = setupSchema(cfg)
	if err != nil {
		return err
	}

	go func() error {
		err = fetchResources(context.Background(), clients, regions)
		if err != nil {
			return err
		}
		return nil
	}()

	go checkUpgrade()

	err = runServer(port, noTracking)
	if err != nil {
		return err
	}

	return nil
}

func runServer(port int, noTracking bool) error {
	log.Infof("Komiser version: %s, commit: %s, buildt: %s", Version, Commit, Buildtime)

	r := v1.Endpoints(context.Background(), noTracking, db)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"profile", "X-Requested-With", "Content-Type", "Authorization"},
	})
	loggedRouter := handlers.LoggingHandler(os.Stdout, cors.Handler(r))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter)
	if err != nil {
		return err
	} else {
		log.Info("Server started on port %d", port)
	}

	return nil
}

func setupSchema(c *models.Config) error {
	var sqldb *sql.DB
	var err error

	if len(c.SQLite.File) > 0 {
		sqldb, err = sql.Open(sqliteshim.ShimName, fmt.Sprintf("file:%s?cache=shared", c.SQLite.File))
		if err != nil {
			return err
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

	_, err = db.NewCreateTable().Model((*models.Resource)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	_, err = db.NewCreateTable().Model((*models.View)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func fetchResources(ctx context.Context, clients []providers.ProviderClient, regions []string) error {
	for _, client := range clients {
		if client.AWSClient != nil {
			go func(ctx context.Context, client providers.ProviderClient, regions []string) {
				aws.FetchResources(ctx, client, regions, db)
			}(ctx, client, regions)
		} else if client.DigitalOceanClient != nil {
			go func(ctx context.Context, client providers.ProviderClient) {
				do.FetchResources(ctx, client, db)
			}(ctx, client)
		} else if client.OciClient != nil {
			go func(ctx context.Context, client providers.ProviderClient) {
				oci.FetchResources(ctx, client, db)
			}(ctx, client)
		}
	}
	return nil
}

func checkUpgrade() {
	url := "https://api.github.com/repos/tailwarden/komiser/releases/latest"
	type GHRelease struct {
		Version string `json:"tag_name"`
	}

	var myClient = &http.Client{Timeout: 5 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		log.Warnf("Failed to check for new version: %s", err)
		return
	}
	defer r.Body.Close()

	target := new(GHRelease)
	err = json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		log.Warnf("Failed to decode new release version: %s", err)
		return
	}

	v1, err := version.NewVersion(Version)
	if err != nil {
		log.Warnf("Failed to parse version: %s", err)
	} else {
		v2, err := version.NewVersion(target.Version)
		if err != nil {
			log.Warnf("Failed to parse version: %s", err)
		} else {
			if v1.LessThan(v2) {
				log.Warnf("Newer Komiser version is available: %s", target.Version)
				log.Warnf("Upgrade instructions: https://github.com/tailwarden/komiser")
			}
		}
	}
}
