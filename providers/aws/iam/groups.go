package iam

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func Groups(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config iam.ListGroupsInput
	iamClient := iam.NewFromConfig(*client.AWSClient)

	output, err := iamClient.ListGroups(ctx, &config)
	if err != nil {
		return resources, err
	}

	for _, group := range output.Groups {
		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "IAM Group",
			ResourceId: *group.Arn,
			Region:     client.AWSClient.Region,
			Name:       *group.GroupName,
			Cost:       0,
			CreatedAt:  *group.CreateDate,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/iamv2/home?region=%s#/groups/details/%s", client.AWSClient.Region, client.AWSClient.Region, *group.GroupName),
		})

		config.Marker = output.Marker
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "IAM Group",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
