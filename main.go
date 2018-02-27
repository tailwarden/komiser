package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	. "./models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/elbv2"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	cache "github.com/patrickmn/go-cache"
)

var (
	cfg         aws.Config
	memoryCache *cache.Cache
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func EC2RegionHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("ec2_region")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeInstancesPerRegion(cfg)
		memoryCache.Set("ec2_region", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func EC2FamilyHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("ec2_family")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeInstancesPerFamily(cfg)
		memoryCache.Set("ec2_family", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func EC2StateHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("ec2_state")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeInstancesPerState(cfg)
		memoryCache.Set("ec2_state", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func EBSSizeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("ebs_size")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeVolumesTotalSize(cfg)
		memoryCache.Set("ebs_size", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func EBSFamilyHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("ebs_family")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeVolumesPerFamily(cfg)
		memoryCache.Set("ebs_family", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func EBSStateHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("ebs_state")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeVolumesPerState(cfg)
		memoryCache.Set("ebs_state", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func VPCTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("vpc_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeVPCsTotal(cfg)
		memoryCache.Set("vpc_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func ACLTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("acl_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeACLsTotal(cfg)
		memoryCache.Set("acl_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func SecurityGroupTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("sg_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeSecurityGroupsTotal(cfg)
		memoryCache.Set("sg_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func NatGatewayTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("nat_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeNatGatewaysTotal(cfg)
		memoryCache.Set("nat_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func ElasticIPTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("eip_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeElasticIPsTotal(cfg)
		memoryCache.Set("eip_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func InternetGatewayTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("igw_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeInternetGatewaysTotal(cfg)
		memoryCache.Set("igw_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func RouteTableTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("rt_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeRouteTablesTotal(cfg)
		memoryCache.Set("rt_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func KeyPairTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("kp_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeKeyPairsTotal(cfg)
		memoryCache.Set("kp_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func AutoScalingGroupTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("asg_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeAutoScalingGroupsTotal(cfg)
		memoryCache.Set("asg_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func ElasticLoadBalancerFamilyHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("elb_family")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeElasticLoadBalancerPerFamily(cfg)
		memoryCache.Set("elb_family", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func S3TotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("s3_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeS3BucketsTotal(cfg)
		memoryCache.Set("s3_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func CostAndUsageHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("cost_usage")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeCostAndUsage(cfg)
		memoryCache.Set("cost_usage", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func LambdaPerRuntimeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("lambda_runtime")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeLambdaFunctionsPerRuntime(cfg)
		memoryCache.Set("lambda_runtime", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func RDSInstancePerEngineHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("rds_engine")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeRDSInstancesPerEngine(cfg)
		memoryCache.Set("rds_engine", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func DynamoDBTableTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("dynamodb_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeDynamoDBTablesTotal(cfg)
		memoryCache.Set("dynamodb_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func DynamoDBProvisionedThroughputHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("dynamodb_throughput")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeDynamoDBTablesProvisionedThroughput(cfg)
		memoryCache.Set("dynamodb_throughput", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func SnapshotTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("snapshot_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeSnapshotsTotal(cfg)
		memoryCache.Set("snapshot_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func SnapshotSizeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := memoryCache.Get("snapshot_size")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := describeSnapshotsSize(cfg)
		memoryCache.Set("snapshot_size", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func init() {
	memoryCache = cache.New(10*time.Minute, 10*time.Minute)
	cfg, _ = external.LoadDefaultAWSConfig()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ec2/region", EC2RegionHandler) //OK
	r.HandleFunc("/ec2/family", EC2FamilyHandler) //Ok
	r.HandleFunc("/ec2/state", EC2StateHandler)   //OK
	r.HandleFunc("/ebs/size", EBSSizeHandler)     //OK
	r.HandleFunc("/ebs/family", EBSFamilyHandler) //OK
	r.HandleFunc("/ebs/state", EBSStateHandler)
	r.HandleFunc("/vpc/total", VPCTotalHandler)                            //OK
	r.HandleFunc("/acl/total", ACLTotalHandler)                            //OK
	r.HandleFunc("/security_group/total", SecurityGroupTotalHandler)       //OK
	r.HandleFunc("/nat/total", NatGatewayTotalHandler)                     //OK
	r.HandleFunc("/eip/total", ElasticIPTotalHandler)                      //OK
	r.HandleFunc("/key_pair/total", KeyPairTotalHandler)                   //OK
	r.HandleFunc("/route_table/total", RouteTableTotalHandler)             //OK
	r.HandleFunc("/internet_gateway/total", InternetGatewayTotalHandler)   //OK
	r.HandleFunc("/autoscaling_group/total", AutoScalingGroupTotalHandler) //OK
	r.HandleFunc("/elb/family", ElasticLoadBalancerFamilyHandler)
	r.HandleFunc("/s3/total", S3TotalHandler)
	r.HandleFunc("/cost", CostAndUsageHandler) //OK
	r.HandleFunc("/lambda/runtime", LambdaPerRuntimeHandler)
	r.HandleFunc("/rds/engine", RDSInstancePerEngineHandler)
	r.HandleFunc("/dynamodb/total", DynamoDBTableTotalHandler)                 //OK
	r.HandleFunc("/dynamodb/throughput", DynamoDBProvisionedThroughputHandler) //OK
	r.HandleFunc("/snapshot/total", SnapshotTotalHandler)                      //OK
	r.HandleFunc("/snapshot/size", SnapshotSizeHandler)                        //OK
	http.ListenAndServe(":3000", handlers.CORS()(r))
}

func describeInstancesPerRegion(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range getRegions(cfg) {
		instances := getInstances(cfg, region.Name)
		output[region.Name] = len(instances)
	}
	return output
}

func describeInstancesPerState(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range getRegions(cfg) {
		instances := getInstances(cfg, region.Name)
		for _, instance := range instances {
			output[instance.State]++
		}
	}
	return output
}

func describeInstancesPerFamily(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range getRegions(cfg) {
		instances := getInstances(cfg, region.Name)
		for _, instance := range instances {
			output[instance.InstanceType]++
		}
	}
	return output
}

func getInstances(cfg aws.Config, region string) []EC2 {
	cfg.Region = region
	ec2Svc := ec2.New(cfg)
	params := &ec2.DescribeInstancesInput{}
	req := ec2Svc.DescribeInstancesRequest(params)
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfInstances := make([]EC2, 0)
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instanceType, _ := instance.InstanceType.MarshalValue()
			instanceState, _ := instance.State.Name.MarshalValue()
			instanceTags := make([]string, 0)
			for _, tag := range instance.Tags {
				instanceTags = append(instanceTags, *tag.Value)
			}
			listOfInstances = append(listOfInstances, EC2{
				ID:           *instance.InstanceId,
				InstanceType: instanceType,
				LaunchTime:   *instance.LaunchTime,
				Tags:         instanceTags,
				State:        instanceState,
			})
		}
	}
	return listOfInstances
}

func describeVolumesTotalSize(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		volumes := getVolumes(cfg, region.Name)
		for _, volume := range volumes {
			sum += volume.Size
		}
	}
	return sum
}

func describeVolumesPerFamily(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range getRegions(cfg) {
		volumes := getVolumes(cfg, region.Name)
		for _, volume := range volumes {
			output[volume.VolumeType]++
		}
	}
	return output
}

func describeVolumesPerState(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range getRegions(cfg) {
		volumes := getVolumes(cfg, region.Name)
		for _, volume := range volumes {
			output[volume.State]++
		}
	}
	return output
}

func getVolumes(cfg aws.Config, region string) []Volume {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeVolumesRequest(&ec2.DescribeVolumesInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfVolumes := make([]Volume, 0)
	for _, volume := range result.Volumes {
		volumeType, _ := volume.VolumeType.MarshalValue()
		volumeState, _ := volume.State.MarshalValue()
		listOfVolumes = append(listOfVolumes, Volume{
			ID:         *volume.VolumeId,
			AZ:         *volume.AvailabilityZone,
			LaunchTime: *volume.CreateTime,
			Size:       *volume.Size,
			State:      volumeState,
			VolumeType: volumeType,
		})
	}
	return listOfVolumes
}

func describeVPCsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		vpcs := getVPCs(cfg, region.Name)
		sum += int64(len(vpcs))
	}
	return sum
}

func getVPCs(cfg aws.Config, region string) []VPC {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeVpcsRequest(&ec2.DescribeVpcsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfVPCs := make([]VPC, 0)
	for _, vpc := range result.Vpcs {
		vpcState, _ := vpc.State.MarshalValue()
		vpcTags := make([]string, 0)
		for _, tag := range vpc.Tags {
			vpcTags = append(vpcTags, *tag.Value)
		}
		listOfVPCs = append(listOfVPCs, VPC{
			ID:        *vpc.VpcId,
			State:     vpcState,
			CidrBlock: *vpc.CidrBlock,
			Tags:      vpcTags,
		})
	}
	return listOfVPCs
}

func describeACLsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		acls := getNetworkACLs(cfg, region.Name)
		sum += int64(len(acls))
	}
	return sum
}

func getNetworkACLs(cfg aws.Config, region string) []NetworkACL {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeNetworkAclsRequest(&ec2.DescribeNetworkAclsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfNetworkACLs := make([]NetworkACL, 0)
	for _, networkACL := range result.NetworkAcls {
		aclTags := make([]string, 0)
		for _, tag := range networkACL.Tags {
			aclTags = append(aclTags, *tag.Value)
		}
		listOfNetworkACLs = append(listOfNetworkACLs, NetworkACL{
			ID:   *networkACL.NetworkAclId,
			Tags: aclTags,
		})
	}
	return listOfNetworkACLs
}

func describeSecurityGroupsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		sgs := getSecurityGroups(cfg, region.Name)
		sum += int64(len(sgs))
	}
	return sum
}

func getSecurityGroups(cfg aws.Config, region string) []SecurityGroup {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeSecurityGroupsRequest(&ec2.DescribeSecurityGroupsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfSecurityGroups := make([]SecurityGroup, 0)
	for _, securityGroup := range result.SecurityGroups {
		sgTags := make([]string, 0)
		for _, tag := range securityGroup.Tags {
			sgTags = append(sgTags, *tag.Value)
		}
		listOfSecurityGroups = append(listOfSecurityGroups, SecurityGroup{
			Tags: sgTags,
		})
	}
	return listOfSecurityGroups
}

func describeNatGatewaysTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		ngws := getNatGateways(cfg, region.Name)
		sum += int64(len(ngws))
	}
	return sum
}

func getNatGateways(cfg aws.Config, region string) []NatGateway {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeNatGatewaysRequest(&ec2.DescribeNatGatewaysInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfNatGateways := make([]NatGateway, 0)
	for _, ngw := range result.NatGateways {
		ngwState, _ := ngw.State.MarshalValue()
		ngwTags := make([]string, 0)
		for _, tag := range ngw.Tags {
			ngwTags = append(ngwTags, *tag.Value)
		}
		listOfNatGateways = append(listOfNatGateways, NatGateway{
			ID:    *ngw.NatGatewayId,
			State: ngwState,
			Tags:  ngwTags,
		})
	}
	return listOfNatGateways
}

func describeElasticIPsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		ips := getElasticIPs(cfg, region.Name)
		sum += int64(len(ips))
	}
	return sum
}

func getElasticIPs(cfg aws.Config, region string) []EIP {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeAddressesRequest(&ec2.DescribeAddressesInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfElasticIPs := make([]EIP, 0)
	for _, address := range result.Addresses {
		addressTags := make([]string, 0)
		for _, tag := range address.Tags {
			addressTags = append(addressTags, *tag.Value)
		}
		listOfElasticIPs = append(listOfElasticIPs, EIP{
			PublicIP: *address.PublicIp,
			Tags:     addressTags,
		})
	}
	return listOfElasticIPs
}

func describeInternetGatewaysTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		igws := getInternetGateways(cfg, region.Name)
		sum += int64(len(igws))
	}
	return sum
}

func getInternetGateways(cfg aws.Config, region string) []InternetGateway {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeInternetGatewaysRequest(&ec2.DescribeInternetGatewaysInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfInternetGateways := make([]InternetGateway, 0)
	for _, igw := range result.InternetGateways {
		igwTags := make([]string, 0)
		for _, tag := range igw.Tags {
			igwTags = append(igwTags, *tag.Value)
		}
		listOfInternetGateways = append(listOfInternetGateways, InternetGateway{
			ID:   *igw.InternetGatewayId,
			Tags: igwTags,
		})
	}
	return listOfInternetGateways
}

func describeRouteTablesTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		rts := getRouteTables(cfg, region.Name)
		sum += int64(len(rts))
	}
	return sum
}

func getRouteTables(cfg aws.Config, region string) []RouteTable {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeRouteTablesRequest(&ec2.DescribeRouteTablesInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfRouteTables := make([]RouteTable, 0)
	for _, rt := range result.RouteTables {
		rtTags := make([]string, 0)
		for _, tag := range rt.Tags {
			rtTags = append(rtTags, *tag.Value)
		}
		listOfRouteTables = append(listOfRouteTables, RouteTable{
			ID:   *rt.RouteTableId,
			Tags: rtTags,
		})
	}
	return listOfRouteTables
}

func describeKeyPairsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		kps := getKeyPairs(cfg, region.Name)
		sum += int64(len(kps))
	}
	return sum
}

func getKeyPairs(cfg aws.Config, region string) []KeyPair {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeKeyPairsRequest(&ec2.DescribeKeyPairsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfKeyPairs := make([]KeyPair, 0)
	for _, kp := range result.KeyPairs {
		listOfKeyPairs = append(listOfKeyPairs, KeyPair{
			KeyName: *kp.KeyName,
		})
	}
	return listOfKeyPairs
}

func describeAutoScalingGroupsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		asgs := getAutoScalingGroups(cfg, region.Name)
		sum += int64(len(asgs))
	}
	return sum
}

func getAutoScalingGroups(cfg aws.Config, region string) []AutoScaling {
	cfg.Region = region
	svc := autoscaling.New(cfg)
	req := svc.DescribeAutoScalingGroupsRequest(&autoscaling.DescribeAutoScalingGroupsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfAutoScalingGroups := make([]AutoScaling, 0)
	for _, asg := range result.AutoScalingGroups {
		asgTags := make([]string, 0)
		for _, tag := range asg.Tags {
			asgTags = append(asgTags, *tag.Value)
		}
		listOfAutoScalingGroups = append(listOfAutoScalingGroups, AutoScaling{
			Status: *asg.Status,
			Tags:   asgTags,
			ARN:    *asg.AutoScalingGroupARN,
		})
	}
	return listOfAutoScalingGroups
}

func describeElasticLoadBalancerPerFamily(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range getRegions(cfg) {
		elbs := getElasticLoadBalancers(cfg, region.Name)
		for _, elb := range elbs {
			output[elb.Type]++
		}
	}
	return output
}

func getElasticLoadBalancers(cfg aws.Config, region string) []LoadBalancer {
	cfg.Region = region
	svc := elbv2.New(cfg)
	req := svc.DescribeLoadBalancersRequest(&elbv2.DescribeLoadBalancersInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfElasticLoadBalancers := make([]LoadBalancer, 0)
	for _, lb := range result.LoadBalancers {
		lbType, _ := lb.Type.MarshalValue()
		listOfElasticLoadBalancers = append(listOfElasticLoadBalancers, LoadBalancer{
			DNSName: *lb.DNSName,
			State:   lb.State.String(),
			Type:    lbType,
		})
	}
	return listOfElasticLoadBalancers
}

func describeS3BucketsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		buckets := getS3Buckets(cfg, region.Name)
		sum += int64(len(buckets))
	}
	return sum
}

/*
func describeS3BucketsSize(cfg aws.Config) {
	svc := s3.New(cfg)
	req := svc.ListObjectsRequest(&s3.ListObjectsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	result.Name
}*/

func getS3Buckets(cfg aws.Config, region string) []Bucket {
	cfg.Region = region
	svc := s3.New(cfg)
	req := svc.ListBucketsRequest(&s3.ListBucketsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfBuckets := make([]Bucket, 0)
	for _, bucket := range result.Buckets {
		listOfBuckets = append(listOfBuckets, Bucket{
			Name: *bucket.Name,
		})
	}
	return listOfBuckets
}

func describeCostAndUsage(cfg aws.Config) []Cost {
	currentTime := time.Now().Local()
	start := currentTime.AddDate(0, -6, 0).Format("2006-01-02")
	end := currentTime.Format("2006-01-02")
	svc := costexplorer.New(cfg)
	req := svc.GetCostAndUsageRequest(&costexplorer.GetCostAndUsageInput{
		Metrics:     []string{"BlendedCost"},
		Granularity: costexplorer.GranularityMonthly,
		TimePeriod: &costexplorer.DateInterval{
			Start: &start,
			End:   &end,
		},
	})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	costs := make([]Cost, 0)
	for _, res := range result.ResultsByTime {
		start, _ := time.Parse("2006-01-02", *res.TimePeriod.Start)
		end, _ := time.Parse("2006-01-02", *res.TimePeriod.End)
		amount, _ := strconv.ParseFloat(*res.Total["BlendedCost"].Amount, 64)
		costs = append(costs, Cost{
			Start:  start,
			End:    end,
			Amount: amount,
			Unit:   *res.Total["BlendedCost"].Unit,
		})
	}
	return costs
}

func describeLambdaFunctionsPerRuntime(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range getRegions(cfg) {
		functions := getLambdaFunctions(cfg, region.Name)
		for _, lambda := range functions {
			output[lambda.Runtime]++
		}
	}
	return output
}

func getLambdaFunctions(cfg aws.Config, region string) []Lambda {
	cfg.Region = region
	svc := lambda.New(cfg)
	req := svc.ListFunctionsRequest(&lambda.ListFunctionsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfFunctions := make([]Lambda, 0)
	for _, lambda := range result.Functions {
		runtime, _ := lambda.Runtime.MarshalValue()
		listOfFunctions = append(listOfFunctions, Lambda{
			Name:    *lambda.FunctionName,
			Memory:  *lambda.MemorySize,
			Runtime: runtime,
		})
	}
	return listOfFunctions
}

func describeRDSInstancesPerEngine(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range getRegions(cfg) {
		instances := getRDSInstances(cfg, region.Name)
		for _, instance := range instances {
			output[instance.Engine]++
		}
	}
	return output
}

func getRDSInstances(cfg aws.Config, region string) []DBInstance {
	cfg.Region = region
	svc := rds.New(cfg)
	req := svc.DescribeDBInstancesRequest(&rds.DescribeDBInstancesInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfInstances := make([]DBInstance, 0)
	for _, instance := range result.DBInstances {
		listOfInstances = append(listOfInstances, DBInstance{
			Status:           *instance.DBInstanceStatus,
			StorageType:      *instance.StorageType,
			AllocatedStorage: *instance.AllocatedStorage,
			InstanceClass:    *instance.DBInstanceClass,
			Engine:           *instance.Engine,
		})
	}
	return listOfInstances
}

func describeDynamoDBTablesTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		tables := getDynamoDBTables(cfg, region.Name)
		sum += int64(len(tables))
	}
	return sum
}

func describeDynamoDBTablesProvisionedThroughput(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range getRegions(cfg) {
		for _, table := range getDynamoDBTables(cfg, region.Name) {
			cfg.Region = region.Name
			svc := dynamodb.New(cfg)
			req := svc.DescribeTableRequest(&dynamodb.DescribeTableInput{
				TableName: &table.Name,
			})
			result, err := req.Send()
			if err != nil {
				log.Fatal(err)
			}
			output["readCapacity"] += int(*result.Table.ProvisionedThroughput.ReadCapacityUnits)
			output["writeCapacity"] += int(*result.Table.ProvisionedThroughput.WriteCapacityUnits)
		}
	}
	return output
}

func getDynamoDBTables(cfg aws.Config, region string) []Table {
	cfg.Region = region
	svc := dynamodb.New(cfg)
	req := svc.ListTablesRequest(&dynamodb.ListTablesInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfTables := make([]Table, 0)
	for _, table := range result.TableNames {
		listOfTables = append(listOfTables, Table{
			Name: table,
		})
	}
	return listOfTables
}

func describeCloudWatchAlarmsPerState(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range getRegions(cfg) {
		alarms := getCloudWatchAlarms(cfg, region.Name)
		for _, alarm := range alarms {
			output[alarm.State]++
		}
	}
	return output
}

func getCloudWatchAlarms(cfg aws.Config, region string) []Alarm {
	cfg.Region = region
	svc := cloudwatch.New(cfg)
	req := svc.DescribeAlarmsRequest(&cloudwatch.DescribeAlarmsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfAlarms := make([]Alarm, 0)
	for _, alarm := range result.MetricAlarms {
		alarmState, _ := alarm.StateValue.MarshalValue()
		listOfAlarms = append(listOfAlarms, Alarm{
			Name:  *alarm.AlarmName,
			State: alarmState,
		})
	}
	return listOfAlarms
}

func describeSnapshotsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		snapshots := getSnapshots(cfg, region.Name)
		sum += int64(len(snapshots))
	}
	return sum
}

func describeSnapshotsSize(cfg aws.Config) int64 {
	var sum int64
	for _, region := range getRegions(cfg) {
		snapshots := getSnapshots(cfg, region.Name)
		for _, snapshot := range snapshots {
			sum += snapshot.VolumeSize
		}
	}
	return sum
}

func getSnapshots(cfg aws.Config, region string) []Snapshot {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeSnapshotsRequest(&ec2.DescribeSnapshotsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfSnapshots := make([]Snapshot, 0)
	for _, snapshot := range result.Snapshots {
		snapshotState, _ := snapshot.State.MarshalValue()
		listOfSnapshots = append(listOfSnapshots, Snapshot{
			State:      snapshotState,
			VolumeSize: *snapshot.VolumeSize,
		})
	}
	return listOfSnapshots
}

func getRegions(cfg aws.Config) []Region {
	svc := ec2.New(cfg)
	req := svc.DescribeRegionsRequest(&ec2.DescribeRegionsInput{})
	regions, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfRegions := make([]Region, 0, len(regions.Regions))
	for _, region := range regions.Regions {
		listOfRegions = append(listOfRegions, Region{
			Name: *region.RegionName,
		})
	}
	return listOfRegions
}
