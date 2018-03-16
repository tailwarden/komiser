package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeInstancesPerRegion(cfg aws.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		instances, err := aws.getInstances(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		output[region.Name] = len(instances)
	}
	return output, nil
}

func (aws AWS) DescribeInstancesPerState(cfg aws.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		instances, err := aws.getInstances(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		for _, instance := range instances {
			output[instance.State]++
		}
	}
	return output, nil
}

func (aws AWS) DescribeInstancesPerFamily(cfg aws.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		instances, err := aws.getInstances(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		for _, instance := range instances {
			output[instance.InstanceType]++
		}
	}
	return output, nil
}

func (aws AWS) getInstances(cfg aws.Config, region string) ([]EC2, error) {
	cfg.Region = region
	ec2Svc := ec2.New(cfg)
	params := &ec2.DescribeInstancesInput{}
	req := ec2Svc.DescribeInstancesRequest(params)
	result, err := req.Send()
	if err != nil {
		return []EC2{}, err
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
	return listOfInstances, nil
}
