package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "github.com/mlabouardy/komiser/handlers/aws"
	. "github.com/mlabouardy/komiser/handlers/digitalocean"
	. "github.com/mlabouardy/komiser/handlers/gcp"
	. "github.com/mlabouardy/komiser/handlers/integrations"
	. "github.com/mlabouardy/komiser/handlers/settings"
	. "github.com/mlabouardy/komiser/services/cache"
	"github.com/robfig/cron"
	"github.com/rs/cors"

	//	. "github.com/mlabouardy/komiser/services/ini"
	. "github.com/mlabouardy/komiser/handlers/azure"
	"github.com/urfave/cli"
)

const (
	DEFAULT_PORT           = 3000
	DEFAULT_DURATION       = 30
	DEFAULT_ALERT_SCHEDULE = "0 9 * * * *"
)

func startServer(port int, cache Cache, dataset string, multiple bool, schedule string, regions []string) {
	cache.Connect()

	services := make(map[string]interface{}, 0)

	digitaloceanHandler := NewDigitalOceanHandler(cache)
	gcpHandler := NewGCPHandler(cache, dataset)
	awsHandler := NewAWSHandler(cache, multiple, regions, services)
	azureHandler := NewAzureHandler(cache)
	alertHandler := NewAlertHandler(awsHandler, gcpHandler, azureHandler)
	accountHandler := NewAccountHandler(cache, awsHandler, gcpHandler, azureHandler, digitaloceanHandler, services)
	c := cron.New()
	c.AddFunc(schedule, alertHandler.DailyNotifHandler)
	c.Start()

	r := mux.NewRouter()

	r.HandleFunc("/accounts", accountHandler.ListCloudAccountsHandler)
	r.HandleFunc("/regions", accountHandler.ListActiveRegionsHandler)
	r.HandleFunc("/billing/providers", accountHandler.CostBreakdownByCloudProviderHandler)
	r.HandleFunc("/billing/accounts", accountHandler.CostBreakdownByCloudAccountHandler)
	r.HandleFunc("/billing/regions", accountHandler.CostBreakdownByCloudRegionHandler)
	r.HandleFunc("/views", accountHandler.ListViewsHandler).Methods("GET")
	r.HandleFunc("/views", accountHandler.NewViewHandler).Methods("POST")

	// AWS supported services
	r.HandleFunc("/aws/ec2/regions", awsHandler.EC2InstancesHandler)
	r.HandleFunc("/aws/lambda/functions", awsHandler.LambdaFunctionHandler)
	r.HandleFunc("/aws/s3/buckets", awsHandler.S3BucketsHandler)
	r.HandleFunc("/aws/dynamodb/tables", awsHandler.DynamoDBTableHandler)
	r.HandleFunc("/aws/vpc", awsHandler.VPCHandler)
	r.HandleFunc("/aws/route_tables", awsHandler.RouteTableHandler)
	r.HandleFunc("/aws/security_groups", awsHandler.SecurityGroupHandler)
	r.HandleFunc("/aws/sqs/queues", awsHandler.SQSQueuesHandler)
	r.HandleFunc("/aws/ecs", awsHandler.ECSHandler)
	r.HandleFunc("/aws/vpc/subnets", awsHandler.DescribeSubnetsHandler)

	// DigitalOcean supported services
	r.HandleFunc("/digitalocean/droplets", digitaloceanHandler.DropletsHandler)
	r.HandleFunc("/digitalocean/snapshots", digitaloceanHandler.SnapshotsHandler)
	r.HandleFunc("/digitalocean/volumes", digitaloceanHandler.VolumesHandler)
	r.HandleFunc("/digitalocean/databases", digitaloceanHandler.DatabasesHandler)

	// GCP supported services
	r.HandleFunc("/gcp/compute/instances", gcpHandler.ComputeInstancesHandler)

	// Deprecated
	r.HandleFunc("/integrations", alertHandler.ListIntegrationsHandler)
	r.HandleFunc("/integrations/slack", alertHandler.SetupSlackHandler).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(assetFS()))

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"profile", "X-Requested-With", "Content-Type", "Authorization"},
	})
	loggedRouter := handlers.LoggingHandler(os.Stdout, cors.Handler(r))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter)
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
				cli.IntFlag{
					Name:  "duration, d",
					Usage: "Cache expiration time",
					Value: DEFAULT_DURATION,
				},
				cli.StringFlag{
					Name:  "redis, r",
					Usage: "Redis server",
				},
				cli.StringFlag{
					Name:  "dataset, ds",
					Usage: "BigQuery Bill dataset",
				},
				cli.StringFlag{
					Name:  "regions, re",
					Usage: "Restrict Komiser inspection to list of regions",
				},
				cli.StringFlag{
					Name:  "cron, c",
					Usage: "Daily budget alert schedule",
					Value: DEFAULT_ALERT_SCHEDULE,
				},
				cli.BoolFlag{
					Name:  "multiple, m",
					Usage: "Enable multiple AWS accounts",
				},
			},
			Action: func(c *cli.Context) error {
				port := c.Int("port")
				duration := c.Int("duration")
				redis := c.String("redis")
				dataset := c.String("dataset")
				multiple := c.Bool("multiple")
				schedule := c.String("cron")
				regions := c.String("regions")

				listOfRegions := []string{}

				if len(regions) > 0 {
					listOfRegions = strings.Split(regions, ",")
					log.Println("Restrict Komiser inspection to the following AWS regions:", listOfRegions)
				}

				var cache Cache

				if port == 0 {
					port = DEFAULT_PORT
				}
				if duration == 0 {
					duration = DEFAULT_DURATION
				}

				if redis == "" {
					cache = &Memory{
						Expiration: time.Duration(duration),
					}
				} else {
					cache = &Redis{
						Addr:       redis,
						Expiration: time.Duration(duration),
					}
				}

				startServer(port, cache, dataset, multiple, schedule, listOfRegions)
				return nil
			},
		},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Command not found %q !", command)
	}
	app.Run(os.Args)
}
