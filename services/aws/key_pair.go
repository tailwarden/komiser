package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeKeyPairsTotal(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		kps, err := aws.getKeyPairs(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(kps))
	}
	return sum, nil
}

func (aws AWS) getKeyPairs(cfg aws.Config, region string) ([]KeyPair, error) {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeKeyPairsRequest(&ec2.DescribeKeyPairsInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return []KeyPair{}, err
	}
	listOfKeyPairs := make([]KeyPair, 0)
	for _, kp := range result.KeyPairs {
		listOfKeyPairs = append(listOfKeyPairs, KeyPair{
			KeyName: *kp.KeyName,
		})
	}
	return listOfKeyPairs, nil
}
