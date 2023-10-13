package kinesis

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

type KinesisClient interface {
	ListStreamConsumers(ctx context.Context, params *kinesis.ListStreamConsumersInput, optFns ...func(*kinesis.Options)) (*kinesis.ListStreamConsumersOutput, error)
}

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
			consumers, err := getStreamConsumers(ctx, kinesisClient, stream, client.Name, client.AWSClient.Region)
			if err != nil {
				return resources, err
			}
			resources = append(resources, consumers...)
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

func getStreamConsumers(ctx context.Context, kinesisClient KinesisClient, stream types.StreamSummary, clientName, region string) ([]Resource, error) {
	resources := make([]Resource, 0)
	config := kinesis.ListStreamConsumersInput{
		StreamARN: aws.String(aws.ToString(stream.StreamARN)),
	}

	for {
		output, err := kinesisClient.ListStreamConsumers(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, consumer := range output.Consumers {
			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    clientName,
				Service:    "Kinesis EFO Consumer",
				ResourceId: *consumer.ConsumerARN,
				Region:     region,
				Name:       *consumer.ConsumerName,
				Cost:       0,
				CreatedAt:  *consumer.ConsumerCreationTimestamp,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/kinesis/home?region=%s#/streams/details/%s/registeredConsumers/%s", region, region, aws.ToString(stream.StreamName), *consumer.ConsumerName),
			})
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return resources, nil
}
