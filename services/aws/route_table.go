package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeRouteTablesTotal(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		rts, err := aws.getRouteTables(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(rts))
	}
	return sum, nil
}

func (aws AWS) getRouteTables(cfg aws.Config, region string) ([]RouteTable, error) {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeRouteTablesRequest(&ec2.DescribeRouteTablesInput{})
	result, err := req.Send(context.Background())
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
