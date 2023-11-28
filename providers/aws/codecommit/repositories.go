package codecommit

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Repositories(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var listRepoParams codecommit.ListRepositoriesInput
	resources := make([]models.Resource, 0)
	codecommitClient := codecommit.NewFromConfig(*client.AWSClient)

	for {
		output, err := codecommitClient.ListRepositories(ctx, &listRepoParams)
		if err != nil {
			return resources, err
		}

		for _, repository := range output.Repositories {
			outputTags, err := codecommitClient.ListTagsForResource(ctx, &codecommit.ListTagsForResourceInput{
				ResourceArn: repository.RepositoryId,
			})

			tags := make([]models.Tag, 0)
			if err == nil {
				for _, tag := range outputTags.Tags {
					tags = append(tags, models.Tag{
						Key:   tag,
						Value: outputTags.Tags[tag],
					})
				}
			}

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "CodeCommit",
				ResourceId: *repository.RepositoryId,
				Region:     client.AWSClient.Region,
				Name:       *repository.RepositoryName,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/codesuite/codecommit/repositories/%s/browse?region=%s", client.AWSClient.Region, *repository.RepositoryName, client.AWSClient.Region),
			})
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}

		listRepoParams.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "CodeCommit",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
