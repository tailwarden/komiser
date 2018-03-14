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
	r.HandleFunc("/ec2/region", awsHandler.EC2RegionHandler)
	r.HandleFunc("/ec2/family", awsHandler.EC2FamilyHandler)
	r.HandleFunc("/ec2/state", awsHandler.EC2StateHandler)
	r.HandleFunc("/ebs/size", awsHandler.EBSSizeHandler)
	r.HandleFunc("/ebs/family", awsHandler.EBSFamilyHandler)
	r.HandleFunc("/vpc/total", awsHandler.VPCTotalHandler)
	r.HandleFunc("/acl/total", awsHandler.ACLTotalHandler)
	r.HandleFunc("/security_group/total", awsHandler.SecurityGroupTotalHandler)
	r.HandleFunc("/nat/total", awsHandler.NatGatewayTotalHandler)
	r.HandleFunc("/eip/total", awsHandler.ElasticIPTotalHandler)
	r.HandleFunc("/key_pair/total", awsHandler.KeyPairTotalHandler)
	r.HandleFunc("/route_table/total", awsHandler.RouteTableTotalHandler)
	r.HandleFunc("/internet_gateway/total", awsHandler.InternetGatewayTotalHandler)
	r.HandleFunc("/autoscaling_group/total", awsHandler.AutoScalingGroupTotalHandler)
	r.HandleFunc("/elb/family", awsHandler.ElasticLoadBalancerFamilyHandler)
	r.HandleFunc("/s3/total", awsHandler.S3TotalHandler)
	r.HandleFunc("/cost", awsHandler.CostAndUsageHandler)
	r.HandleFunc("/lambda/runtime", awsHandler.LambdaPerRuntimeHandler)
	r.HandleFunc("/rds/engine", awsHandler.RDSInstancePerEngineHandler)
	r.HandleFunc("/dynamodb/total", awsHandler.DynamoDBTableTotalHandler)
	r.HandleFunc("/dynamodb/throughput", awsHandler.DynamoDBProvisionedThroughputHandler)
	r.HandleFunc("/snapshot/total", awsHandler.SnapshotTotalHandler)
	r.HandleFunc("/snapshot/size", awsHandler.SnapshotSizeHandler)
	r.HandleFunc("/sqs/total", awsHandler.SQSTotalHandler)
	r.HandleFunc("/sns/total", awsHandler.TopicsTotalHandler)
	r.HandleFunc("/hosted_zone/total", awsHandler.HostedZoneTotalHandler)
	r.HandleFunc("/role/total", awsHandler.IAMRolesTotalHandler)
	r.HandleFunc("/group/total", awsHandler.IAMGroupsTotalHandler)
	r.HandleFunc("/user/total", awsHandler.IAMUsersTotalHandler)
	r.HandleFunc("/policy/total", awsHandler.IAMPoliciesTotalHandler)
	r.HandleFunc("/cloudwatch/state", awsHandler.CloudWatchAlarmsStateHandler)
	r.HandleFunc("/cloudfront/total", awsHandler.CloudFrontDistributionsTotalHandler)
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
				},
				cli.IntFlag{
					Name:  "duration, d",
					Usage: "Cache expiration time",
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
