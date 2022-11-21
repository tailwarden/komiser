package sqs

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func Queues(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config sqs.ListQueuesInput
	sqsClient := sqs.NewFromConfig(*client.AWSClient)

	for {
		output, err := sqsClient.ListQueues(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, queue := range output.QueueUrls {
			outputTags, err := sqsClient.ListQueueTags(ctx, &sqs.ListQueueTagsInput{
				QueueUrl: &queue,
			})

			tags := make([]Tag, 0)

			if err == nil {
				for key, value := range outputTags.Tags {
					tags = append(tags, Tag{
						Key:   key,
						Value: value,
					})
				}
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "SQS",
				ResourceId: queue,
				Region:     client.AWSClient.Region,
				Name:       queue,
				Cost:       0,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/sqs/v2/home?region=%s#/queues/%s", client.AWSClient.Region, client.AWSClient.Region, queue),
			})
		}

		if aws.ToString(config.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "SQS",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
