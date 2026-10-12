package codeguru

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func ReviewerRepositoryAssociations(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var listAssociationsParams codegurureviewer.ListRepositoryAssociationsInput
	resources := make([]models.Resource, 0)
	codeguruClient := codegurureviewer.NewFromConfig(*client.AWSClient)

	for {
		output, err := codeguruClient.ListRepositoryAssociations(ctx, &listAssociationsParams)
		if err != nil {
			return resources, err
		}

		for _, association := range output.RepositoryAssociationSummaries {
			tags := make([]models.Tag, 0)
			if association.AssociationArn != nil {
				outputTags, err := codeguruClient.ListTagsForResource(ctx, &codegurureviewer.ListTagsForResourceInput{
					ResourceArn: association.AssociationArn,
				})
				if err == nil {
					for key, value := range outputTags.Tags {
						tags = append(tags, models.Tag{
							Key:   key,
							Value: value,
						})
					}
				}
			}

			lastUpdated := time.Now()
			if association.LastUpdatedTimeStamp != nil {
				lastUpdated = *association.LastUpdatedTimeStamp
			}

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "CodeGuru",
				ResourceId: aws.ToString(association.AssociationArn),
				Region:     client.AWSClient.Region,
				Name:       aws.ToString(association.Name),
				CreatedAt:  lastUpdated,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Metadata: map[string]string{
					"owner":         aws.ToString(association.Owner),
					"provider":      string(association.ProviderType),
					"state":         string(association.State),
					"associationId": aws.ToString(association.AssociationId),
				},
				Link: fmt.Sprintf("https://%s.console.aws.amazon.com/codeguru/reviewer?region=%s#/associations", client.AWSClient.Region, client.AWSClient.Region),
			})
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}

		listAssociationsParams.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "CodeGuru",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
