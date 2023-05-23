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

func Streams(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config kinesis.ListStreamsInput
	kinesisClient := kinesis.NewFromConfig(*client.AWSClient)

	for {
		output, err := kinesisClient.ListStreams(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, stream := range output.StreamSummaries {
			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Kinesis Stream",
				ResourceId: *stream.StreamARN,
				Region:     client.AWSClient.Region,
				Name:       *stream.StreamName,
				Cost:       0,
				CreatedAt:  *stream.StreamCreationTimestamp,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/kinesis/home?region=%s#/streams/details/%s", client.AWSClient.Region, client.AWSClient.Region, *stream.StreamName),
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
		"service":   "Kinesis Stream",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
