package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeInternetGatewaysTotal(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		igws, err := aws.getInternetGateways(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(igws))
	}
	return sum, nil
}

func (aws AWS) getInternetGateways(cfg awsConfig.Config, region string) ([]InternetGateway, error) {
	cfg.Region = region
	svc := ec2.NewFromConfig(cfg)
	result, err := svc.DescribeInternetGateways(context.Background(), &ec2.DescribeInternetGatewaysInput{})
	if err != nil {
		return []InternetGateway{}, err
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
	return listOfInternetGateways, nil
}
