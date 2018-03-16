package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (aws AWS) DescribeS3Buckets(cfg aws.Config) (int, error) {
	svc := s3.New(cfg)
	req := svc.ListBucketsRequest(&s3.ListBucketsInput{})
	result, err := req.Send()
	if err != nil {
		return 0, err
	}
	return len(result.Buckets), nil
}
