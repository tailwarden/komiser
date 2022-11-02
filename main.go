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

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/digitalocean/godo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "github.com/mlabouardy/komiser/handlers"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers/aws"
	. "github.com/mlabouardy/komiser/providers/digitalocean"
	"github.com/rs/cors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/urfave/cli"
	"gopkg.in/ini.v1"
)

const (
	DEFAULT_PORT = 3000
)

func startServer(port int) {

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("KOMISER_POSTGRES_URI"))))
	db := bun.NewDB(sqldb, pgdialect.New())

	db.NewCreateTable().Model((*Resource)(nil)).Exec(context.Background())

	// if AWS is enabled
	f, err := ini.Load(config.DefaultSharedCredentialsFilename())
	for _, section := range f.Sections() {
		profileName := strings.ToLower(section.Name())
		log.Println("Fetch resources from AWS:", profileName)
		cfg, _ := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(profileName))
		go func() {
			FetchAwsData(context.Background(), cfg, profileName, db)
		}()
	}
	// if DigitalOcean is supported
	if os.Getenv("KOMISER_DIGITALOCEAN_TOKEN") != "" {
		digitalOceanClient := godo.NewFromToken(os.Getenv("KOMISER_DIGITALOCEAN_TOKEN"))
		go func() {
			FetchDigitalOceanData(context.Background(), digitalOceanClient, "Default", db)
		}()
	}

	r := mux.NewRouter()

	resourcesHandler := NewResourcesHandler(context.Background(), db)

	r.HandleFunc("/resources", resourcesHandler.ListResourcesHandler)
	r.HandleFunc("/regions", resourcesHandler.RegionsCounterHandler)
	r.HandleFunc("/resources/count", resourcesHandler.ResourcesCounterHandler)

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
