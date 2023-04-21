package iam

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func Policies(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	configPolicies := iam.ListPoliciesInput{
		Scope: types.PolicyScopeTypeLocal,
	}
	var configTags iam.ListPolicyTagsInput
	iamClient := iam.NewFromConfig(*client.AWSClient)

	for {
		outputPolicies, err := iamClient.ListPolicies(ctx, &configPolicies)
		if err != nil {
			return resources, err
		}

		for _, policy := range outputPolicies.Policies {
			tags := make([]Tag, 0)
			for {
				configTags.PolicyArn = policy.Arn
				outputPolicyTags, err := iamClient.ListPolicyTags(ctx, &configTags)
				if err != nil {
					return resources, err
				}

				for _, t := range outputPolicyTags.Tags {
					tags = append(tags, Tag{
						Key:   *t.Key,
						Value: *t.Value,
					})
				}

				if aws.ToString(outputPolicyTags.Marker) == "" {
					break
				}

				configTags.Marker = outputPolicyTags.Marker
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "IAM Policy",
				ResourceId: *policy.Arn,
				Region:     client.AWSClient.Region,
				Name:       *policy.PolicyName,
				Cost:       0,
				CreatedAt:  *policy.CreateDate,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/iam/home#/policies/%s", client.AWSClient.Region, *policy.Arn),
			})
		}

		if aws.ToString(outputPolicies.Marker) == "" {
			break
		}

		configPolicies.Marker = outputPolicies.Marker
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "IAM Policy",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
