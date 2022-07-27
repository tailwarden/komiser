package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

func (aws AWS) ListStreams(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := kinesis.NewFromConfig(cfg)
		res, err := svc.ListStreams(context.Background(), &kinesis.ListStreamsInput{})
		if err != nil {
			return sum, err
		}

		sum += int64(len(res.StreamNames))
	}
	return sum, nil
}

func (awsClient AWS) ListShards(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := kinesis.NewFromConfig(cfg)
		res, err := svc.ListStreams(context.Background(), &kinesis.ListStreamsInput{})
		if err != nil {
			return sum, err
		}

		for _, stream := range res.StreamNames {
			res, err := svc.ListShards(context.Background(), &kinesis.ListShardsInput{
				StreamName: &stream,
			})
			if err != nil {
				return sum, err
			}

			sum += int64(len(res.Shards))
		}
	}
	return sum, nil
}
