package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

func (aws AWS) DescribeCloudFrontDistributionsTotal(cfg aws.Config) (int, error) {
	svc := cloudfront.New(cfg)
	req := svc.ListDistributionsRequest(&cloudfront.ListDistributionsInput{})
	result, err := req.Send()
	if err != nil {
		return 0, err
	}
	return len(result.DistributionList.Items), nil
}
