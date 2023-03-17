package iam

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func InstanceProfiles(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config iam.ListInstanceProfilesInput
	iamClient := iam.NewFromConfig(*client.AWSClient)

	output, err := iamClient.ListInstanceProfiles(ctx, &config)
	if err != nil {
		return resources, err
	}

	for _, instanceprofile := range output.InstanceProfiles {
		outputTags, err := iamClient.ListInstanceProfileTags(ctx, &iam.ListInstanceProfileTagsInput{
			InstanceProfileName: instanceprofile.InstanceProfileName,
		})
		if err != nil {
			return resources, err
		}

		tags := make([]Tag, 0)
		for _, t := range outputTags.Tags {
			tags = append(tags, Tag{
				Key:   *t.Key,
				Value: *t.Value,
			})
		}

		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "IAM Instance Profile",
			ResourceId: *instanceprofile.Arn,
			Region:     client.AWSClient.Region,
			Name:       *instanceprofile.InstanceProfileName,
			Cost:       0,
			CreatedAt:  *instanceprofile.CreateDate,
			Tags:       tags,
			FetchedAt:  time.Now(),
		})

		config.Marker = output.Marker
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "IAM Instance Profile",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
