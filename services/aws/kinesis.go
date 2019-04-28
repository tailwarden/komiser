package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

func (aws AWS) ListStreams(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := kinesis.New(cfg)
		req := svc.ListStreamsRequest(&kinesis.ListStreamsInput{})
		res, err := req.Send(context.Background())
		if err != nil {
			return sum, err
		}

		sum += int64(len(res.StreamNames))
	}
	return sum, nil
}

func (awsClient AWS) ListShards(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := kinesis.New(cfg)
		req := svc.ListStreamsRequest(&kinesis.ListStreamsInput{})
		res, err := req.Send(context.Background())
		if err != nil {
			return sum, err
		}

		for _, stream := range res.StreamNames {
			req := svc.ListShardsRequest(&kinesis.ListShardsInput{
				StreamName: aws.String(stream),
			})
			res, err := req.Send(context.Background())
			if err != nil {
				return sum, err
			}

			sum += int64(len(res.Shards))
		}
	}
	return sum, nil
}
