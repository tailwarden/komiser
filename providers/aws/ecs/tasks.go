package instances

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func Tasks(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecs.ListTasksInput
	ecsClient := ecs.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	for {
		output, err := ecsClient.ListTasks(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, task := range output.TaskArns {
			resourceArn := fmt.Sprintf("arn:aws:ecs:%s:%s:task/%s", client.AWSClient.Region, *accountId, task)

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "ECS Task",
				ResourceId: resourceArn,
				Region:     client.AWSClient.Region,
				Name:       task,
				Cost:       0,
				FetchedAt:  time.Now(),
			})
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.Printf("[%s] Fetched %d AWS ECS tasks from %s\n", client.Name, len(resources), client.AWSClient.Region)
	return resources, nil
}
