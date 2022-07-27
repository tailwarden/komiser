package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

func (aws AWS) DescribeHostedZones(cfg awsConfig.Config) (int, error) {
	cfg.Region = "us-east-1"
	svc := route53.NewFromConfig(cfg)
	result, err := svc.ListHostedZones(context.Background(), &route53.ListHostedZonesInput{})
	if err != nil {
		return 0, err
	}
	return len(result.HostedZones), nil
}

func (aws AWS) DescribeARecords(cfg awsConfig.Config) (int64, error) {
	var sum int64
	cfg.Region = "us-east-1"
	svc := route53.NewFromConfig(cfg)
	result, err := svc.ListHostedZones(context.Background(), &route53.ListHostedZonesInput{})
	if err != nil {
		return sum, err
	}
	for _, zone := range result.HostedZones {
		res, err := svc.ListResourceRecordSets(context.Background(), &route53.ListResourceRecordSetsInput{
			HostedZoneId: zone.Id,
		})
		if err != nil {
			return sum, err
		}
		for _, record := range res.ResourceRecordSets {
			if record.Type == types.RRTypeA {
				sum++
			}
		}
	}
	return sum, nil
}
