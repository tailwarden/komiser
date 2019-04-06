package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeInstances(cfg aws.Config) (map[string]interface{}, error) {
	outputInstancesPerRegion := make(map[string]int, 0)
	outputInstancesPerState := make(map[string]int, 0)
	outputInstancesPerFamily := make(map[string]int, 0)
	totalPublicInstances := 0
	totalPrivateInstances := 0
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]interface{}{}, err
	}
	for _, region := range regions {
		instances, err := aws.getInstances(cfg, region.Name)
		if err != nil {
			return map[string]interface{}{}, err
		}
		for _, instance := range instances {
			outputInstancesPerState[instance.State]++
			outputInstancesPerFamily[instance.InstanceType]++
			if instance.Public {
				totalPublicInstances++
			} else {
				totalPrivateInstances++
			}
		}
		outputInstancesPerRegion[region.Name] = len(instances)
	}
	return map[string]interface{}{
		"region":  outputInstancesPerRegion,
		"state":   outputInstancesPerState,
		"family":  outputInstancesPerFamily,
		"public":  totalPublicInstances,
		"private": totalPrivateInstances,
	}, nil
}

func (awsClient AWS) getInstances(cfg aws.Config, region string) ([]EC2, error) {
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
			isPublic := true
			if instance.PublicIpAddress == aws.String("") {
				isPublic = false
			}
			listOfInstances = append(listOfInstances, EC2{
				ID:           *instance.InstanceId,
				InstanceType: instanceType,
				LaunchTime:   *instance.LaunchTime,
				Tags:         instanceTags,
				State:        instanceState,
				Public:       isPublic,
			})
		}
	}
	return listOfInstances, nil
}
