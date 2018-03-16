package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func (aws AWS) DescribeHostedZones(cfg aws.Config) (int, error) {
	svc := route53.New(cfg)
	req := svc.ListHostedZonesRequest(&route53.ListHostedZonesInput{})
	result, err := req.Send()
	if err != nil {
		return 0, err
	}
	return len(result.HostedZones), nil
}
