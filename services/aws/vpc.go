package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeVPCsTotal(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		vpcs, err := aws.getVPCs(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(vpcs))
	}
	return sum, nil
}

func (aws AWS) DescribeSubnets(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return sum, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := ec2.New(cfg)
		req := svc.DescribeSubnetsRequest(&ec2.DescribeSubnetsInput{})
		res, err := req.Send(context.Background())
		if err != nil {
			return sum, err
		}
		sum += int64(len(res.Subnets))
	}
	return sum, nil
}

func (aws AWS) getVPCs(cfg aws.Config, region string) ([]VPC, error) {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeVpcsRequest(&ec2.DescribeVpcsInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return []VPC{}, err
	}
	listOfVPCs := make([]VPC, 0)
	for _, vpc := range result.Vpcs {
		vpcState, _ := vpc.State.MarshalValue()
		vpcTags := make([]string, 0)
		for _, tag := range vpc.Tags {
			vpcTags = append(vpcTags, *tag.Value)
		}
		listOfVPCs = append(listOfVPCs, VPC{
			ID:        *vpc.VpcId,
			State:     vpcState,
			CidrBlock: *vpc.CidrBlock,
			Tags:      vpcTags,
		})
	}
	return listOfVPCs, nil
}
