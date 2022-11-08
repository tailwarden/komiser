package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/digitalocean/godo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "github.com/mlabouardy/komiser/handlers"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
	. "github.com/mlabouardy/komiser/providers/aws"
	. "github.com/mlabouardy/komiser/providers/digitalocean"
	. "github.com/mlabouardy/komiser/providers/oci"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/rs/cors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/urfave/cli"
)

const (
	DEFAULT_PORT = 3000
)

var db *bun.DB

func setupSchema(uri string) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(uri)))
	db = bun.NewDB(sqldb, pgdialect.New())

	db.NewCreateTable().Model((*Resource)(nil)).IfNotExists().Exec(context.Background())
}

func loadCloudAccounts(komiserConfig ConfigFile) {
	if len(komiserConfig.AWS) > 0 {
		for _, account := range komiserConfig.AWS {
			if account.Source == "CREDENTIALS_FILE" {
				go func(account AWSConfig) {
					cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(account.Profile))
					if err != nil {
						log.Fatal(err)
					}
					providerClient := ProviderClient{
						AWSClient: &cfg,
						Name:      account.Name,
					}
					FetchAwsData(context.Background(), providerClient, db)
				}(account)
			} else if account.Source == "ENVIRONMENT_VARIABLES" {
				go func(account AWSConfig) {
					cfg, err := config.LoadDefaultConfig(context.Background())
					if err != nil {
						log.Fatal(err)
					}
					providerClient := ProviderClient{
						AWSClient: &cfg,
						Name:      account.Name,
					}
					FetchAwsData(context.Background(), providerClient, db)
				}(account)
			}
		}
	}

	if len(komiserConfig.DigitalOcean) > 0 {
		for _, account := range komiserConfig.DigitalOcean {
			go func(account DigitalOceanConfig) {
				digitalOceanClient := godo.NewFromToken(account.Token)
				providerClient := ProviderClient{
					DigitalOceanClient: digitalOceanClient,
					Name:               account.Name,
				}
				FetchDigitalOceanData(context.Background(), providerClient, db)
			}(account)
		}
	}

	if len(komiserConfig.Oci) > 0 {
		for _, account := range komiserConfig.Oci {
			if account.Source == "CREDENTIALS_FILE" {
				go func(account OciConfig) {
					configProvider := common.DefaultConfigProvider()
					providerClient := ProviderClient{
						OciClient: configProvider,
						Name:      account.Name,
					}
					FetchOciData(context.Background(), providerClient, db)
				}(account)
			}
		}

	}
}

func startServer(port int) {
	komiserConfig := &ConfigFile{}
	data, _ := os.ReadFile("config.toml")
	err := toml.Unmarshal([]byte(data), komiserConfig)
	if err != nil {
		log.Fatal(err)
	}

	if komiserConfig.Postgres.URI == "" {
		log.Fatalf("Postgres URI is missing")
	}

	setupSchema(komiserConfig.Postgres.URI)
	loadCloudAccounts(*komiserConfig)

	r := mux.NewRouter()

	resourcesHandler := NewResourcesHandler(context.Background(), db)

	r.HandleFunc("/resources", resourcesHandler.ListResourcesHandler)
	r.HandleFunc("/resources/tags", resourcesHandler.ListResourcesHandler).Methods("POST")
	r.HandleFunc("/resources/count", resourcesHandler.ResourcesCounterHandler)
	r.HandleFunc("/resources/{id}/tags", resourcesHandler.UpdateTagsHandler).Methods("POST")
	r.HandleFunc("/regions", resourcesHandler.RegionsCounterHandler)
	r.HandleFunc("/costs", resourcesHandler.CostCounterHandler)
	r.HandleFunc("/stats", resourcesHandler.StatsHandler)

	r.PathPrefix("/").Handler(http.FileServer(assetFS()))

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"profile", "X-Requested-With", "Content-Type", "Authorization"},
	})
	loggedRouter := handlers.LoggingHandler(os.Stdout, cors.Handler(r))
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Server started on port %d", port)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Komiser"
	app.Version = "3.0.0"
	app.Usage = "Cloud Environment Inspector"
	app.Copyright = "Komiser - https://komiser.io"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Mohamed Labouardy",
			Email: "mohamed@oraculi.io",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start server",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port, p",
					Usage: "Server port",
					Value: DEFAULT_PORT,
				},
				cli.StringFlag{
					Name:  "regions, re",
					Usage: "Restrict Komiser inspection to list of regions",
				},
			},
			Action: func(c *cli.Context) error {
				port := c.Int("port")
				regions := c.String("regions")

				listOfRegions := []string{}

				if len(regions) > 0 {
					listOfRegions = strings.Split(regions, ",")
					log.Println("Restrict Komiser inspection to the following AWS regions:", listOfRegions)
				}

				startServer(port)
				return nil
			},
		},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Command not found %q !", command)
	}
	app.Run(os.Args)
}
