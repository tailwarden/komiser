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

func EventSubscription(ctx context.Context, client ProviderClient) ([]Resource, error) {
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

		for _, eventSubscription := range output.EventSubscriptionsList { // TODO: capitalization convention?
			if eventSubscription.CustSubscriptionId != nil { // TODO: is this equivalent to filesystem.Name from the efs example?

				resourceArn := fmt.Sprintf("arn:aws:redshift:%s:%s:eventsubscripion/%s", client.AWSClient.Region, *accountId, *eventSubscription.CustSubscriptionId) // TODO: is this arn format correct
				outputTags, err := redshiftClient.DescribeTags(ctx, &redshift.DescribeTagsInput{
					ResourceName: &resourceArn, // TODO: is ResourceName here equivalent to ResourceId in the efs example?
				})

				tags := make([]Tag, 0)

				if err == nil {
					for _, tag := range outputTags.TaggedResources { // TODO: this is slightly different than in the efs example. Is it correct?
						tags = append(tags, Tag{
							Key:   *tag.Tag.Key,
							Value: *tag.Tag.Value,
						})
					}
				}

				monthlyCost := float64(0) // TODO: what is the monthly cost

				resources = append(resources, Resource{
					Provider:   "AWS",
					Account:    client.Name,
					Service:    "Redshift",
					ResourceId: resourceArn,
					Region:     client.AWSClient.Region,
					Name:       *eventSubscription.CustSubscriptionId,
					Cost:       monthlyCost,
					Tags:       tags,
					FetchedAt:  time.Now(),
					Link:       fmt.Sprintf("https://%s.console.aws.amaxon.com/redshift/home?region=%s/event-subscriptions/%s", client.AWSClient.Region, client.AWSClient.Region, eventSubscription.CustSubscriptionId), // TODO: verify that the link format is correct
				})
			}
		}

		if aws.ToString(output.Marker) == "" { // TODO: is output.Marker here playing the same role as ouput.NextMarker in the efs example?
			break
		}
		config.Marker = output.Marker
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Redshift",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil

}
