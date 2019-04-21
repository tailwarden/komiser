package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeElasticIPsTotal(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		ips, err := aws.getElasticIPs(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(ips))
	}
	return sum, nil
}

func (aws AWS) getElasticIPs(cfg aws.Config, region string) ([]EIP, error) {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeAddressesRequest(&ec2.DescribeAddressesInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return []EIP{}, err
	}
	listOfElasticIPs := make([]EIP, 0)
	for _, address := range result.Addresses {
		if address.AssociationId == nil {
			addressTags := make([]string, 0)
			for _, tag := range address.Tags {
				addressTags = append(addressTags, *tag.Value)
			}
			listOfElasticIPs = append(listOfElasticIPs, EIP{
				PublicIP: *address.PublicIp,
				Tags:     addressTags,
			})
		}
	}
	return listOfElasticIPs, nil
}
