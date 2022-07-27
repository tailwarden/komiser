package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeACLsTotal(cfg awsConfig.Config) (int64, error) {
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

func (aws AWS) getNetworkACLs(cfg awsConfig.Config, region string) ([]NetworkACL, error) {
	cfg.Region = region
	svc := ec2.NewFromConfig(cfg)
	result, err := svc.DescribeNetworkAcls(context.Background(), &ec2.DescribeNetworkAclsInput{})
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
