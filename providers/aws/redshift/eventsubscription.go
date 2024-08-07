package redshift

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func EventSubscriptions(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	var config redshift.DescribeEventSubscriptionsInput
	redshiftClient := redshift.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "Redshift")
	if err != nil {
		log.Warnln("Couldn't fetch Redshift cost and usage:", err)
	}

	for {
		output, err := redshiftClient.DescribeEventSubscriptions(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, eventSubscription := range output.EventSubscriptionsList {
			if eventSubscription.CustSubscriptionId != nil {

				resourceArn := fmt.Sprintf("arn:aws:redshift:%s:%s:eventsubscripion/%s", client.AWSClient.Region, *accountId, *eventSubscription.CustSubscriptionId)
				outputTags := eventSubscription.Tags

				tags := make([]models.Tag, 0)

				for _, tag := range outputTags {
					tags = append(tags, models.Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}

				monthlyCost := float64(0)

				resources = append(resources, models.Resource{
					Provider:   "AWS",
					Account:    client.Name,
					Service:    "Redshift EventSubscription",
					ResourceId: resourceArn,
					Region:     client.AWSClient.Region,
					Name:       *eventSubscription.CustSubscriptionId,
					Cost:       monthlyCost,
					Metadata: map[string]string{
						"serviceCost": fmt.Sprint(serviceCost),
					},
					Tags:      tags,
					FetchedAt: time.Now(),
					Link:      fmt.Sprintf("https://%s.console.aws.amaxon.com/redshift/home?region=%s/event-subscriptions/%s", client.AWSClient.Region, client.AWSClient.Region, *eventSubscription.CustSubscriptionId),
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
