package iam

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

const (
	iamUserLinkTemplate = "https://%s.console.aws.amazon.com/iamv2/home?region=%s#/users/details/%s"
	awsProvider         = "AWS"
	iamUserService      = "IAM User"
)

func Users(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var resources []models.Resource

	iamClient := iam.NewFromConfig(*client.AWSClient)

	paginator := iam.NewListUsersPaginator(iamClient, &iam.ListUsersInput{})

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "IAM")
	if err != nil {
		log.Warnln("Couldn't fetch IAM cost and usage:", err)
	}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			log.WithError(err).Error("Failed to list IAM users")
			return nil, fmt.Errorf("failed to list IAM users: %w", err)
		}

		for _, o := range output.Users {
			var tags []models.Tag
			for _, t := range o.Tags {
				tags = append(tags, models.Tag{
					Key:   aws.ToString(t.Key),
					Value: aws.ToString(t.Value),
				})
			}

			resources = append(resources, models.Resource{
				Provider:   awsProvider,
				Account:    client.Name,
				Service:    iamUserService,
				ResourceId: aws.ToString(o.Arn),
				Region:     client.AWSClient.Region,
				Name:       aws.ToString(o.UserName),
				Cost:       0,
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				CreatedAt:  *o.CreateDate,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf(iamUserLinkTemplate, client.AWSClient.Region, client.AWSClient.Region, aws.ToString(o.UserName)),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  awsProvider,
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   iamUserService,
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
