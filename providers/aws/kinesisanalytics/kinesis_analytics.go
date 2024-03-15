package kinesisanalytics

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func KinesisAnalytics(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	var config kinesisanalyticsv2.ListApplicationsInput
	kinesisAnalyticsClient := kinesisanalyticsv2.NewFromConfig(*client.AWSClient)
	for {
		output, err := kinesisAnalyticsClient.ListApplications(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, app := range output.ApplicationSummaries {
			outputTags, err := kinesisAnalyticsClient.ListTagsForResource(ctx, &kinesisanalyticsv2.ListTagsForResourceInput{
				ResourceARN: app.ApplicationARN,
			})

			tags := make([]models.Tag, 0)

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
				Service:    "Kinesis Analytics Application",
				ResourceId: *app.ApplicationARN,
				Region:     client.AWSClient.Region,
				Name:       *app.ApplicationName,
				Cost:       0,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/eks/home?region=%s#/clusters", client.AWSClient.Region, client.AWSClient.Region),
			})
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":    "AWS",
		"account":     client.Name,
		"region":      client.AWSClient.Region,
		"service":     "Kinesis Analytics Application",
		"resources":   len(resources),
		"serviceCost": fmt.Sprint(0),
	}).Info("Fetched resources")
	return resources, nil
}
