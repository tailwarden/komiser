package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func (aws AWS) DescribeHostedZones(cfg aws.Config) (int, error) {
	svc := route53.New(cfg)
	req := svc.ListHostedZonesRequest(&route53.ListHostedZonesInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return 0, err
	}
	return len(result.HostedZones), nil
}

func (aws AWS) DescribeARecords(cfg aws.Config) (int64, error) {
	var sum int64
	svc := route53.New(cfg)
	req := svc.ListHostedZonesRequest(&route53.ListHostedZonesInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return sum, err
	}
	for _, zone := range result.HostedZones {
		req := svc.ListResourceRecordSetsRequest(&route53.ListResourceRecordSetsInput{
			HostedZoneId: zone.Id,
		})
		res, err := req.Send(context.Background())
		if err != nil {
			return sum, err
		}
		for _, record := range res.ResourceRecordSets {
			if record.Type == route53.RRTypeA {
				sum++
			}
		}
	}
	return sum, nil
}
