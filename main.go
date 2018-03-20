package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "github.com/mlabouardy/komiser/handlers"
	cache "github.com/patrickmn/go-cache"
	"github.com/urfave/cli"
)

const (
	DEFAULT_PORT     = 3000
	DEFAULT_DURATION = 30
)

func startServer(port int, duration int) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	}

	cache := cache.New(time.Duration(duration)*time.Minute, time.Duration(duration)*time.Minute)
	awsHandler := NewAWSHandler(cfg, cache)

	r := mux.NewRouter()
	r.HandleFunc("/ec2", awsHandler.EC2InstancesHandler)
	r.HandleFunc("/ebs", awsHandler.EBSHandler)
	r.HandleFunc("/vpc", awsHandler.VPCHandler)
	r.HandleFunc("/acl", awsHandler.ACLHandler)
	r.HandleFunc("/security_group", awsHandler.SecurityGroupHandler)
	r.HandleFunc("/nat", awsHandler.NatGatewayHandler)
	r.HandleFunc("/key_pair", awsHandler.KeyPairHandler)
	r.HandleFunc("/route_table", awsHandler.RouteTableHandler)
	r.HandleFunc("/internet_gateway", awsHandler.InternetGatewayHandler)
	r.HandleFunc("/eip", awsHandler.ElasticIPHandler)
	r.HandleFunc("/autoscaling_group", awsHandler.AutoScalingGroupHandler)
	r.HandleFunc("/elb", awsHandler.ElasticLoadBalancerHandler)
	r.HandleFunc("/cost", awsHandler.CostAndUsageHandler)
	r.HandleFunc("/lambda", awsHandler.LambdaFunctionHandler)
	r.HandleFunc("/rds", awsHandler.RDSInstanceHandler)
	r.HandleFunc("/dynamodb", awsHandler.DynamoDBTableHandler)
	r.HandleFunc("/snapshot", awsHandler.SnapshotHandler)
	r.HandleFunc("/sqs", awsHandler.SQSQueuesHandler)
	r.HandleFunc("/sns", awsHandler.SNSTopicsHandler)
	r.HandleFunc("/hosted_zone", awsHandler.HostedZoneHandler)
	r.HandleFunc("/iam/role", awsHandler.IAMRolesHandler)
	r.HandleFunc("/iam/group", awsHandler.IAMGroupsHandler)
	r.HandleFunc("/iam/user", awsHandler.IAMUsersHandler)
	r.HandleFunc("/iam/policy", awsHandler.IAMPoliciesHandler)
	r.HandleFunc("/ecs", awsHandler.ECSHandler)
	r.HandleFunc("/cloudwatch", awsHandler.CloudWatchAlarmsHandler)
	r.HandleFunc("/cloudfront", awsHandler.CloudFrontDistributionsHandler)
	r.HandleFunc("/s3", awsHandler.S3BucketsHandler)
	r.PathPrefix("/").Handler(http.FileServer(assetFS()))
	loggedRouter := handlers.LoggingHandler(os.Stdout, handlers.CORS()(r))
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
	app.Version = "1.0.0"
	app.Usage = "AWS Environment Inspector"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Mohamed Labouardy",
			Email: "mohamed@labouardy.com",
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
					Value: 3000,
				},
				cli.IntFlag{
					Name:  "duration, d",
					Usage: "Cache expiration time",
					Value: 30,
				},
			},
			Action: func(c *cli.Context) error {
				port := c.Int("port")
				duration := c.Int("duration")
				if port == 0 {
					port = DEFAULT_PORT
				}
				if duration == 0 {
					duration = DEFAULT_DURATION
				}
				startServer(port, duration)
				return nil
			},
		},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Command not found %q !", command)
	}
	app.Run(os.Args)
}
