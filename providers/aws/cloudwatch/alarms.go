package cloudwatch

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func Alarms(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config cloudwatch.DescribeAlarmsInput
	cloudWatchClient := cloudwatch.NewFromConfig(*client.AWSClient)
	for {
		output, err := cloudWatchClient.DescribeAlarms(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, alarm := range output.MetricAlarms {
			outputTags, err := cloudWatchClient.ListTagsForResource(ctx, &cloudwatch.ListTagsForResourceInput{
				ResourceARN: alarm.AlarmArn,
			})

			tags := make([]Tag, 0)

			if err == nil {
				for _, tag := range outputTags.Tags {
					tags = append(tags, Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "CloudWatch",
				ResourceId: *alarm.AlarmArn,
				Region:     client.AWSClient.Region,
				Name:       *alarm.AlarmName,
				Cost:       0,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/cloudwatch/home?region=%s#alarmsV2:alarm/%s", client.AWSClient.Region, client.AWSClient.Region, *alarm.AlarmName),
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
		"service":   "CloudWatch",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
