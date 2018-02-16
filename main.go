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
	fmt.Println(describeVolumesPerState(cfg))
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
