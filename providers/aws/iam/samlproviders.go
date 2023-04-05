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

func SamlProviders(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var configSamlProviders iam.ListSAMLProvidersInput
	var configTags iam.ListSAMLProviderTagsInput
	iamClient := iam.NewFromConfig(*client.AWSClient)

	outputSamlProviders, err := iamClient.ListSAMLProviders(ctx, &configSamlProviders)
	if err != nil {
		return resources, err
	}

	for _, samlprovider := range outputSamlProviders.SAMLProviderList {
		tags := make([]Tag, 0)
		for {
			configTags.SAMLProviderArn = samlprovider.Arn
			outputSamlProviderTags, err := iamClient.ListSAMLProviderTags(ctx, &configTags)
			if err != nil {
				return resources, err
			}

			for _, t := range outputSamlProviderTags.Tags {
				tags = append(tags, Tag{
					Key:   *t.Key,
					Value: *t.Value,
				})
			}

			if aws.ToString(outputSamlProviderTags.Marker) == "" {
				break
			}

			configTags.Marker = outputSamlProviderTags.Marker
		}

		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "IAM SAML Provider",
			ResourceId: *samlprovider.Arn,
			Region:     client.AWSClient.Region,
			Name:       *samlprovider.Arn,
			Cost:       0,
			CreatedAt:  *samlprovider.CreateDate,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/iamv2/home?region=%s#/identity_providers/details/%s", client.AWSClient.Region, client.AWSClient.Region, *samlprovider.Arn),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "IAM SAML Provider",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
