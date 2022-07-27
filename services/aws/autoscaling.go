package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeAutoScalingGroups(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		asgs, err := aws.getAutoScalingGroups(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(asgs))
	}
	return sum, nil
}

func (aws AWS) getAutoScalingGroups(cfg awsConfig.Config, region string) ([]AutoScaling, error) {
	cfg.Region = region
	svc := autoscaling.NewFromConfig(cfg)
	result, err := svc.DescribeAutoScalingGroups(context.Background(), &autoscaling.DescribeAutoScalingGroupsInput{})
	if err != nil {
		return []AutoScaling{}, err
	}
	listOfAutoScalingGroups := make([]AutoScaling, 0)
	for _, asg := range result.AutoScalingGroups {
		asgTags := make([]string, 0)
		for _, tag := range asg.Tags {
			asgTags = append(asgTags, *tag.Value)
		}
		listOfAutoScalingGroups = append(listOfAutoScalingGroups, AutoScaling{
			ARN: *asg.AutoScalingGroupARN,
		})
	}
	return listOfAutoScalingGroups, nil
}
