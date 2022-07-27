package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (awsClient AWS) DescribeElasticIPsTotal(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		ips, err := awsClient.getElasticIPs(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(ips))
	}
	return sum, nil
}

func (awsClient AWS) getElasticIPs(cfg awsConfig.Config, region string) ([]EIP, error) {
	cfg.Region = region
	svc := ec2.NewFromConfig(cfg)
	result, err := svc.DescribeAddresses(context.Background(), &ec2.DescribeAddressesInput{})
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
