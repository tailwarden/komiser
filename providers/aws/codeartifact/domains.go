package codeartifact

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Domains(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var listDomainsParams codeartifact.ListDomainsInput
	resources := make([]models.Resource, 0)
	codeartifactClient := codeartifact.NewFromConfig(*client.AWSClient)

	for {
		output, err := codeartifactClient.ListDomains(ctx, &listDomainsParams)
		if err != nil {
			return resources, err
		}

		for _, domain := range output.Domains {
			outputTags, err := codeartifactClient.ListTagsForResource(ctx, &codeartifact.ListTagsForResourceInput{
				ResourceArn: domain.Arn,
			})

			tags := make([]models.Tag, 0)
			if err == nil {
				for _, tag := range outputTags.Tags {
					tags = append(tags, models.Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}
			}

			createdAt := time.Now()
			if domain.CreatedTime != nil {
				createdAt = *domain.CreatedTime
			}

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "CodeArtifact",
				ResourceId: *domain.Arn,
				Region:     client.AWSClient.Region,
				Name:       *domain.Name,
				CreatedAt:  createdAt,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Metadata: map[string]string{
					"owner":  *domain.Owner,
					"status": string(domain.Status),
				},
				Link: fmt.Sprintf("https://%s.console.aws.amazon.com/codesuite/codeartifact/domains/%s?region=%s", client.AWSClient.Region, *domain.Name, client.AWSClient.Region),
			})
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}

		listDomainsParams.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "CodeArtifact",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
