package cloudwatch

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func MetricStreams(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	cloudWatchMetricsClient := cloudwatch.NewFromConfig(*client.AWSClient)

	input := &cloudwatch.ListMetricStreamsInput{}
	for {
		output, err := cloudWatchMetricsClient.ListMetricStreams(ctx, input)
		if err != nil {
			return resources, err
		}

		for _, stream := range output.Entries {
			tags := make([]models.Tag, 0)

			streamArn := aws.ToString(stream.Arn)
			tagInput := &cloudwatch.ListTagsForResourceInput{
				ResourceARN: &streamArn,
			}

			tagOutput, err := cloudWatchMetricsClient.ListTagsForResource(ctx, tagInput)
			if err == nil {
				for _, tag := range tagOutput.Tags {
					tags = append(tags, models.Tag{
						Key:   aws.ToString(tag.Key),
						Value: aws.ToString(tag.Value),
					})
				}
			}

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "CloudWatch Metric Stream",
				ResourceId: streamArn,
				Region:     client.AWSClient.Region,
				Name:       aws.ToString(stream.Name),
				Cost:       0,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/cloudwatch/home?region=%s#metric-streams:streamsList/%s", client.AWSClient.Region, client.AWSClient.Region, aws.ToString(stream.Name)),
			})
		}

		if output.NextToken == nil {
			break
		}

		input.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "CloudWatch Metric Stream",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
