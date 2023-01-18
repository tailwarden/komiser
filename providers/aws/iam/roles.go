package iam

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func Roles(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config iam.ListRolesInput
	iamClient := iam.NewFromConfig(*client.AWSClient)
	output, err := iamClient.ListRoles(context.Background(), &config)
	if err != nil {
		return resources, err
	}

	for _, o := range output.Roles {
		tags := make([]Tag, 0)

		for _, t := range o.Tags {
			tags = append(tags, Tag{
				Key:   *t.Key,
				Value: *t.Value,
			})
		}

		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "IAM Role",
			ResourceId: *o.Arn,
			Region:     client.AWSClient.Region,
			Name:       *o.RoleName,
			Cost:       0,
			CreatedAt:  *o.CreateDate,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/iamv2/home?region=%s#/roles/details/%s", client.AWSClient.Region, client.AWSClient.Region, *o.RoleName),
		})

		if aws.ToString(output.Marker) == "" {
			break
		}

		config.Marker = output.Marker
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "IAM Role",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
