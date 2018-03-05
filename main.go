package main

import (
	"net/http"
	"time"

	. "./handlers"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	cache "github.com/patrickmn/go-cache"
)

var (
	awsHandler *AWSHandler
)

func init() {
	cache := cache.New(30*time.Minute, 30*time.Minute)
	cfg, _ := external.LoadDefaultAWSConfig()
	awsHandler = NewAWSHandler(cfg, cache)
}

func main() {
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
	http.ListenAndServe(":3000", handlers.CORS()(r))
}
