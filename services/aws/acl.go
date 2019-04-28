package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeACLsTotal(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		acls, err := aws.getNetworkACLs(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(acls))
	}
	return sum, nil
}

func (aws AWS) getNetworkACLs(cfg aws.Config, region string) ([]NetworkACL, error) {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeNetworkAclsRequest(&ec2.DescribeNetworkAclsInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return []NetworkACL{}, err
	}
	listOfNetworkACLs := make([]NetworkACL, 0)
	for _, networkACL := range result.NetworkAcls {
		aclTags := make([]string, 0)
		for _, tag := range networkACL.Tags {
			aclTags = append(aclTags, *tag.Value)
		}
		listOfNetworkACLs = append(listOfNetworkACLs, NetworkACL{
			ID:   *networkACL.NetworkAclId,
			Tags: aclTags,
		})
	}
	return listOfNetworkACLs, nil
}
