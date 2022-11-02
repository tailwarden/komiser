package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	. "github.com/mlabouardy/komiser/models"
)

func Tasks(ctx context.Context, cfg aws.Config, account string) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecs.ListTasksInput
	ecsClient := ecs.NewFromConfig(cfg)
	output, err := ecsClient.ListTasks(context.Background(), &config)
	if err != nil {
		return resources, err
	}

	for _, task := range output.TaskArns {
		resources = append(resources, Resource{
			Provider:  "AWS",
			Account:   account,
			Service:   "ECS Task",
			Region:    cfg.Region,
			Name:      task,
			Cost:      0,
			FetchedAt: time.Now(),
		})

		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.Printf("[%s] Fetched %d AWS ECS tasks from %s\n", account, len(resources), cfg.Region)
	return resources, nil
}
