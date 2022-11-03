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
	var config ecr.ListImagesInput
	ecrClient := ecr.NewFromConfig(*client.AWSClient)

	for {
		output, err := ecrClient.ListImages(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, image := range output.ImageIds {
			resources = append(resources, Resource{
				Provider:  "AWS",
				Account:   client.Name,
				Service:   "ECR",
				Region:    client.AWSClient.Region,
				Name:      *image.ImageTag,
				Cost:      0,
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
