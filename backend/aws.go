package backend

import (
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/elb"
	"github.com/aws/aws-sdk-go-v2/service/elbv2"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	. "github.com/mlabouardy/komiser/models"
)

type AWS struct {
}

func (aws AWS) DescribeInstancesPerRegion(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range aws.getRegions(cfg) {
		instances := aws.getInstances(cfg, region.Name)
		output[region.Name] = len(instances)
	}
	return output
}

func (aws AWS) DescribeInstancesPerState(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range aws.getRegions(cfg) {
		instances := aws.getInstances(cfg, region.Name)
		for _, instance := range instances {
			output[instance.State]++
		}
	}
	return output
}

func (aws AWS) DescribeInstancesPerFamily(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range aws.getRegions(cfg) {
		instances := aws.getInstances(cfg, region.Name)
		for _, instance := range instances {
			output[instance.InstanceType]++
		}
	}
	return output
}

func (aws AWS) getInstances(cfg aws.Config, region string) []EC2 {
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

func (aws AWS) DescribeVolumesTotalSize(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		volumes := aws.getVolumes(cfg, region.Name)
		for _, volume := range volumes {
			sum += volume.Size
		}
	}
	return sum
}

func (aws AWS) DescribeVolumesPerFamily(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range aws.getRegions(cfg) {
		volumes := aws.getVolumes(cfg, region.Name)
		for _, volume := range volumes {
			output[volume.VolumeType]++
		}
	}
	return output
}

func (aws AWS) DescribeVolumesPerState(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range aws.getRegions(cfg) {
		volumes := aws.getVolumes(cfg, region.Name)
		for _, volume := range volumes {
			output[volume.State]++
		}
	}
	return output
}

func (aws AWS) getVolumes(cfg aws.Config, region string) []Volume {
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

func (aws AWS) DescribeVPCsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		vpcs := aws.getVPCs(cfg, region.Name)
		sum += int64(len(vpcs))
	}
	return sum
}

func (aws AWS) getVPCs(cfg aws.Config, region string) []VPC {
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

func (aws AWS) DescribeACLsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		acls := aws.getNetworkACLs(cfg, region.Name)
		sum += int64(len(acls))
	}
	return sum
}

func (aws AWS) getNetworkACLs(cfg aws.Config, region string) []NetworkACL {
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

func (aws AWS) DescribeSecurityGroupsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		sgs := aws.getSecurityGroups(cfg, region.Name)
		sum += int64(len(sgs))
	}
	return sum
}

func (aws AWS) getSecurityGroups(cfg aws.Config, region string) []SecurityGroup {
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

func (aws AWS) DescribeNatGatewaysTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		ngws := aws.getNatGateways(cfg, region.Name)
		sum += int64(len(ngws))
	}
	return sum
}

func (aws AWS) getNatGateways(cfg aws.Config, region string) []NatGateway {
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

func (aws AWS) DescribeElasticIPsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		ips := aws.getElasticIPs(cfg, region.Name)
		sum += int64(len(ips))
	}
	return sum
}

func (aws AWS) getElasticIPs(cfg aws.Config, region string) []EIP {
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

func (aws AWS) DescribeInternetGatewaysTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		igws := aws.getInternetGateways(cfg, region.Name)
		sum += int64(len(igws))
	}
	return sum
}

func (aws AWS) getInternetGateways(cfg aws.Config, region string) []InternetGateway {
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

func (aws AWS) DescribeRouteTablesTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		rts := aws.getRouteTables(cfg, region.Name)
		sum += int64(len(rts))
	}
	return sum
}

func (aws AWS) getRouteTables(cfg aws.Config, region string) []RouteTable {
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

func (aws AWS) DescribeKeyPairsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		kps := aws.getKeyPairs(cfg, region.Name)
		sum += int64(len(kps))
	}
	return sum
}

func (aws AWS) getKeyPairs(cfg aws.Config, region string) []KeyPair {
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

func (aws AWS) DescribeAutoScalingGroupsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		asgs := aws.getAutoScalingGroups(cfg, region.Name)
		sum += int64(len(asgs))
	}
	return sum
}

func (aws AWS) getAutoScalingGroups(cfg aws.Config, region string) []AutoScaling {
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

func (aws AWS) DescribeElasticLoadBalancerPerFamily(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range aws.getRegions(cfg) {
		elbsv1 := aws.getClassicElasticLoadBalancers(cfg, region.Name)
		elbsv2 := aws.getElasticLoadBalancersV2(cfg, region.Name)
		for _, elb := range elbsv1 {
			output[elb.Type]++
		}
		for _, elb := range elbsv2 {
			output[elb.Type]++
		}
	}
	return output
}

func (aws AWS) getClassicElasticLoadBalancers(cfg aws.Config, region string) []LoadBalancer {
	cfg.Region = region
	svc := elb.New(cfg)
	req := svc.DescribeLoadBalancersRequest(&elb.DescribeLoadBalancersInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfElasticLoadBalancers := make([]LoadBalancer, 0)
	for _, lb := range result.LoadBalancerDescriptions {
		listOfElasticLoadBalancers = append(listOfElasticLoadBalancers, LoadBalancer{
			DNSName: *lb.DNSName,
			Type:    "classic",
		})
	}
	return listOfElasticLoadBalancers
}

func (aws AWS) getElasticLoadBalancersV2(cfg aws.Config, region string) []LoadBalancer {
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

func (aws AWS) DescribeS3BucketsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		buckets := aws.getS3Buckets(cfg, region.Name)
		sum += int64(len(buckets))
	}
	return sum
}

func (aws AWS) getS3Buckets(cfg aws.Config, region string) []Bucket {
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

func (aws AWS) DescribeCostAndUsage(cfg aws.Config) []Cost {
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

func (aws AWS) DescribeLambdaFunctionsPerRuntime(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range aws.getRegions(cfg) {
		functions := aws.getLambdaFunctions(cfg, region.Name)
		for _, lambda := range functions {
			output[lambda.Runtime]++
		}
	}
	return output
}

func (aws AWS) getLambdaFunctions(cfg aws.Config, region string) []Lambda {
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

func (aws AWS) DescribeRDSInstancesPerEngine(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range aws.getRegions(cfg) {
		instances := aws.getRDSInstances(cfg, region.Name)
		for _, instance := range instances {
			output[instance.Engine]++
		}
	}
	return output
}

func (aws AWS) getRDSInstances(cfg aws.Config, region string) []DBInstance {
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

func (aws AWS) DescribeDynamoDBTablesTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		tables := aws.getDynamoDBTables(cfg, region.Name)
		sum += int64(len(tables))
	}
	return sum
}

func (aws AWS) DescribeDynamoDBTablesProvisionedThroughput(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range aws.getRegions(cfg) {
		for _, table := range aws.getDynamoDBTables(cfg, region.Name) {
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

func (aws AWS) getDynamoDBTables(cfg aws.Config, region string) []Table {
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

func (aws AWS) DescribeCloudWatchAlarmsPerState(cfg aws.Config) map[string]int {
	output := make(map[string]int, 0)
	for _, region := range aws.getRegions(cfg) {
		alarms := aws.getCloudWatchAlarms(cfg, region.Name)
		for _, alarm := range alarms {
			output[alarm.State]++
		}
	}
	return output
}

func (aws AWS) getCloudWatchAlarms(cfg aws.Config, region string) []Alarm {
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

func (aws AWS) DescribeSnapshotsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		snapshots := aws.getSnapshots(cfg, region.Name)
		sum += int64(len(snapshots))
	}
	return sum
}

func (aws AWS) DescribeSnapshotsSize(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		snapshots := aws.getSnapshots(cfg, region.Name)
		for _, snapshot := range snapshots {
			sum += snapshot.VolumeSize
		}
	}
	return sum
}

func (aws AWS) getSnapshots(cfg aws.Config, region string) []Snapshot {
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

func (aws AWS) DescribeQueuesTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		queues := aws.getSQS(cfg, region.Name)
		sum += int64(len(queues))
	}
	return sum
}

func (aws AWS) getSQS(cfg aws.Config, region string) []Queue {
	cfg.Region = region
	svc := sqs.New(cfg)
	req := svc.ListQueuesRequest(&sqs.ListQueuesInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfQueues := make([]Queue, 0, len(result.QueueUrls))
	for _, queue := range result.QueueUrls {
		listOfQueues = append(listOfQueues, Queue{
			Name: queue,
		})
	}
	return listOfQueues
}

func (aws AWS) DescribeSNSTopicsTotal(cfg aws.Config) int64 {
	var sum int64
	for _, region := range aws.getRegions(cfg) {
		topics := aws.getSNSTopics(cfg, region.Name)
		sum += int64(len(topics))
	}
	return sum
}

func (aws AWS) getSNSTopics(cfg aws.Config, region string) []Topic {
	cfg.Region = region
	svc := sns.New(cfg)
	req := svc.ListTopicsRequest(&sns.ListTopicsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfTopics := make([]Topic, 0, len(result.Topics))
	for _, topic := range result.Topics {
		listOfTopics = append(listOfTopics, Topic{
			ARN: *topic.TopicArn,
		})
	}
	return listOfTopics
}

func (aws AWS) DescribeHostedZonesTotal(cfg aws.Config) int {
	svc := route53.New(cfg)
	req := svc.ListHostedZonesRequest(&route53.ListHostedZonesInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	return len(result.HostedZones)
}

func (aws AWS) DescribeIAMRolesTotal(cfg aws.Config) int {
	svc := iam.New(cfg)
	req := svc.ListRolesRequest(&iam.ListRolesInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	return len(result.Roles)
}

func (aws AWS) DescribeIAMUsersTotal(cfg aws.Config) int {
	svc := iam.New(cfg)
	req := svc.ListUsersRequest(&iam.ListUsersInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	return len(result.Users)
}

func (aws AWS) DescribeIAMGroupsTotal(cfg aws.Config) int {
	svc := iam.New(cfg)
	req := svc.ListGroupsRequest(&iam.ListGroupsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	return len(result.Groups)
}

func (aws AWS) DescribeIAMPoliciesTotal(cfg aws.Config) int {
	svc := iam.New(cfg)
	req := svc.ListPoliciesRequest(&iam.ListPoliciesInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	return len(result.Policies)
}

func (aws AWS) DescribeCloudFrontDistributionsTotal(cfg aws.Config) int {
	svc := cloudfront.New(cfg)
	req := svc.ListDistributionsRequest(&cloudfront.ListDistributionsInput{})
	result, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	return len(result.DistributionList.Items)
}

func (aws AWS) getRegions(cfg aws.Config) []Region {
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
