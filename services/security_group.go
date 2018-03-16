package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeSecurityGroupsTotal(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		sgs, err := aws.getSecurityGroups(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(sgs))
	}
	return sum, nil
}

func (aws AWS) getSecurityGroups(cfg aws.Config, region string) ([]SecurityGroup, error) {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeSecurityGroupsRequest(&ec2.DescribeSecurityGroupsInput{})
	result, err := req.Send()
	if err != nil {
		return []SecurityGroup{}, err
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
	return listOfSecurityGroups, nil
}
