package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeNatGatewaysTotal(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		ngws, err := aws.getNatGateways(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(ngws))
	}
	return sum, nil
}

func (aws AWS) getNatGateways(cfg aws.Config, region string) ([]NatGateway, error) {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeNatGatewaysRequest(&ec2.DescribeNatGatewaysInput{})
	result, err := req.Send()
	if err != nil {
		return []NatGateway{}, err
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
	return listOfNatGateways, nil
}
