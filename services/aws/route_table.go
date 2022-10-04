package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

type AWSRouteTable struct {
	Name   string   `json:"name"`
	ID     string   `json:"id"`
	Region string   `json:"region"`
	Tags   []string `json:"tags"`
}

func (aws AWS) DescribeRouteTablesTotal(cfg awsConfig.Config) ([]AWSRouteTable, error) {
	rts := make([]AWSRouteTable, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return rts, err
	}
	for _, region := range regions {
		rtsResp, err := aws.getRouteTables(cfg, region.Name)
		if err != nil {
			return rts, err
		}

		for _, rt := range rtsResp {
			rts = append(rts, AWSRouteTable{
				Name:   rt.ID,
				ID:     rt.ID,
				Region: region.Name,
				Tags:   rt.Tags,
			})
		}
	}
	return rts, nil
}

func (aws AWS) getRouteTables(cfg awsConfig.Config, region string) ([]RouteTable, error) {
	cfg.Region = region
	svc := ec2.NewFromConfig(cfg)
	result, err := svc.DescribeRouteTables(context.Background(), &ec2.DescribeRouteTablesInput{})
	if err != nil {
		return []RouteTable{}, err
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
	return listOfRouteTables, nil
}
