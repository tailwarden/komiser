package firehose

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func DeliveryStreams(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config firehose.ListDeliveryStreamsInput
	resources := make([]models.Resource, 0)
	deliveryStreamsClient := firehose.NewFromConfig(*client.AWSClient)

	for {
		output, err := deliveryStreamsClient.ListDeliveryStreams(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, deliveryStreamName := range output.DeliveryStreamNames {
			tags := make([]models.Tag, 0)

			outputTags, err := deliveryStreamsClient.ListTagsForDeliveryStream(ctx, &firehose.ListTagsForDeliveryStreamInput{
				DeliveryStreamName: &deliveryStreamName,
			})

			if err == nil {
				for _, tag := range outputTags.Tags {
					tags = append(tags, models.Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}
			}
			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Kinesis Firehose delivery stream",
				Region:     client.AWSClient.Region,
				ResourceId: deliveryStreamName,
				Cost:       0,
				Name:       deliveryStreamName,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/firehose/home?region=%s#/details/%s/monitoring", client.AWSClient.Region, client.AWSClient.Region, deliveryStreamName),
			})
		}

		if *output.HasMoreDeliveryStreams {
			break
		}
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Kinesis Firehose delivery stream",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
