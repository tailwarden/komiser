package redshift

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func EventSubscriptions(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config redshift.DescribeEventSubscriptionsInput
	redshiftClient := redshift.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	for {
		output, err := redshiftClient.DescribeEventSubscriptions(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, eventSubscription := range output.EventSubscriptionsList {
			if eventSubscription.CustSubscriptionId != nil {

				resourceArn := fmt.Sprintf("arn:aws:redshift:%s:%s:eventsubscripion/%s", client.AWSClient.Region, *accountId, *eventSubscription.CustSubscriptionId) // TODO: is this arn format correct
				outputTags := eventSubscription.Tags

				tags := make([]Tag, 0)

				if err == nil {
					for _, tag := range outputTags {
						tags = append(tags, Tag{
							Key:   *tag.Key,
							Value: *tag.Value,
						})
					}
				}

				monthlyCost := float64(0) // TODO: what is the monthly cost

				resources = append(resources, Resource{
					Provider:   "AWS",
					Account:    client.Name,
					Service:    "Redshift EventSubscription",
					ResourceId: resourceArn,
					Region:     client.AWSClient.Region,
					Name:       *eventSubscription.CustSubscriptionId,
					Cost:       monthlyCost,
					Tags:       tags,
					FetchedAt:  time.Now(),
					Link:       fmt.Sprintf("https://%s.console.aws.amaxon.com/redshift/home?region=%s/event-subscriptions/%s", client.AWSClient.Region, client.AWSClient.Region, *eventSubscription.CustSubscriptionId), // TODO: verify that the link format is correct
				})
			}
		}

		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Redshift EventSubscription",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil

}
