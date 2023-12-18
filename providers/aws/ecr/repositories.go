package ecr

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func Repositories(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecr.DescribeRepositoriesInput
	ecrClient := ecr.NewFromConfig(*client.AWSClient)
	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "ECR")
	if err != nil {
		log.Warnln("Couldn't fetch ECR cost and usage:", err)
	}
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
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "ECR",
				ResourceId: *repository.RepositoryArn,
				Region:     client.AWSClient.Region,
				Name:       *repository.RepositoryName,
				Cost:       0.10,
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ecr/repositories/%s", client.AWSClient.Region, *repository.RepositoryName),
			})
		}

		if aws.ToString(config.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "ECR",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
