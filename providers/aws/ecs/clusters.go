package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func EcsClusters(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecs.ListClustersInput
	ecsClient := ecs.NewFromConfig(*client.AWSClient)
	output, err := ecsClient.ListClusters(context.Background(), &config)
	if err != nil {
		return resources, err
	}

	for _, cluster := range output.ClusterArns {
		resources = append(resources, Resource{
			Provider:  "AWS",
			Account:   client.Name,
			Service:   "ECS Cluster",
			Region:    client.AWSClient.Region,
			Name:      cluster,
			Cost:      0,
			FetchedAt: time.Now(),
		})

		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.Printf("[%s] Fetched %d AWS ECS clusters from %s\n", client.Name, len(resources), client.AWSClient.Region)
	return resources, nil
}
