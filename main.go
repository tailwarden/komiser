package main

import (
	"fmt"
	"log"

	. "./models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
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
	fmt.Println(describeNatGatewaysTotal(cfg))
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
