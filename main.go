package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	. "./models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/elbv2"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	//	fmt.Println(describeInstancesPerState(cfg))
	//	fmt.Println(describeInstancesPerRegion(cfg))
	//  fmt.Println(describeInstancesPerFamily(cfg))
	// fmt.Println(describeVolumesPerFamily(cfg))
	// fmt.Println(describeVolumesPerState(cfg))
	// fmt.Println(describeVolumesTotalSize(cfg))
	// fmt.Println(describeACLsTotal(cfg))
	// fmt.Println(describeNatGatewaysTotal(cfg))
	// fmt.Println(describeCostAndUsage(cfg))
	fmt.Println(describeRDSInstancesPerEngine(cfg))
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

func getDynamoDBTables(cfg aws.Config, region string) {
	cfg.Region = region
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
