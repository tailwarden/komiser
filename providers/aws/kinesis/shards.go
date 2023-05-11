package kinesis

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func Shards(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config kinesis.ListShardsInput
	kinesisClient := kinesis.NewFromConfig(*client.AWSClient)

	for {
		output, err := kinesisClient.ListShards(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, shard := range output.Shards {
			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Kinesis Shard",
				ResourceId: *shard.ShardId,
				Region:     client.AWSClient.Region,
				Name:       *shard.ShardId,
				Cost:       0,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/kinesis/home?region=%s#/streams/details/%s", client.AWSClient.Region, client.AWSClient.Region, *shard.ShardId),
			})
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Kinesis Shard",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
