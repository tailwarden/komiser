package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	. "github.com/mlabouardy/komiser/models"
)

func EcsClusters(ctx context.Context, cfg aws.Config, account string) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecs.ListClustersInput
	ecsClient := ecs.NewFromConfig(cfg)
	output, err := ecsClient.ListClusters(context.Background(), &config)
	if err != nil {
		return resources, err
	}

	for _, cluster := range output.ClusterArns {
		resources = append(resources, Resource{
			Provider:  "AWS",
			Account:   account,
			Service:   "ECS Cluster",
			Region:    cfg.Region,
			Name:      cluster,
			Cost:      0,
			FetchedAt: time.Now(),
		})

		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.Printf("[%s] Fetched %d AWS ECS clusters from %s\n", account, len(resources), cfg.Region)
	return resources, nil
}
