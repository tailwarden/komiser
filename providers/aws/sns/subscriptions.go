package sns

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func Subscriptions(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config sns.ListSubscriptionsInput
	snsClient := sns.NewFromConfig(*client.AWSClient)

	for {
		output, err := snsClient.ListSubscriptions(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, subscriptions := range output.Subscriptions {
			outputTags, err := snsClient.ListTagsForResource(ctx, &sns.ListTagsForResourceInput{
				ResourceArn: subscriptions.SubscriptionArn,
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
				Service:    "SNS Subscription",
				ResourceId: *subscriptions.SubscriptionArn,
				Region:     client.AWSClient.Region,
				Name:       *subscriptions.SubscriptionArn,
				Cost:       0,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/sns/v3/home?region=%s#/topic/%s", client.AWSClient.Region, client.AWSClient.Region),
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
		"service":   "SNS Subscription",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
