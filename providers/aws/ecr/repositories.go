package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func Repositories(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecr.DescribeRepositoriesInput
	ecrClient := ecr.NewFromConfig(*client.AWSClient)
	for {
		output, err := ecrClient.DescribeRepositories(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, repository := range output.Repositories {
			outputTags, err := ecrClient.ListTagsForResource(ctx, &ecr.ListTagsForResourceInput{
				ResourceArn: repository.RepositoryArn,
			})

			tags := make([]Tag, 0)

			if err == nil {
				for _, tag := range outputTags.Tags {
					tags = append(tags, Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}
			}

			resources = append(resources, Resource{
				Provider:  "AWS",
				Account:   client.Name,
				Service:   "ECR",
				Region:    client.AWSClient.Region,
				Name:      *repository.RepositoryName,
				Cost:      0,
				Tags:      tags,
				FetchedAt: time.Now(),
			})
		}

		if aws.ToString(config.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.Printf("[%s] Fetched %d AWS ECR repositories from %s\n", client.Name, len(resources), client.AWSClient.Region)
	return resources, nil
}
