package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
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
	result, err := req.Send(context.Background())
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
			isPublic := false
			if instance.PublicIpAddress != nil {
				isPublic = true
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

func (aws AWS) DescribeScheduledInstances(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := ec2.New(cfg)
		req := svc.DescribeScheduledInstancesRequest(&ec2.DescribeScheduledInstancesInput{})
		res, _ := req.Send(context.Background())

		if res != nil {
			for _, set := range res.ScheduledInstanceSet {
				sum += *set.InstanceCount
			}
		}
	}
	return sum, nil
}

func (aws AWS) DescribeReservedInstances(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := ec2.New(cfg)
		req := svc.DescribeReservedInstancesRequest(&ec2.DescribeReservedInstancesInput{})
		res, err := req.Send(context.Background())
		if err != nil {
			return sum, err
		}

		for _, reservation := range res.ReservedInstances {
			sum += *reservation.InstanceCount
		}
	}
	return sum, nil
}

func (aws AWS) DescribeSpotInstances(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := ec2.New(cfg)
		req := svc.DescribeSpotFleetRequestsRequest(&ec2.DescribeSpotFleetRequestsInput{})
		res, err := req.Send(context.Background())
		if err != nil {
			return sum, err
		}

		for _, request := range res.SpotFleetRequestConfigs {
			req2 := svc.DescribeSpotFleetInstancesRequest(&ec2.DescribeSpotFleetInstancesInput{
				SpotFleetRequestId: request.SpotFleetRequestId,
			})
			res2, err := req2.Send(context.Background())
			if err != nil {
				return sum, err
			}

			sum += int64(len(res2.ActiveInstances))
		}
	}
	return sum, nil
}
