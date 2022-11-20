package ecs

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

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
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ecs/home?#/clusters/tasks/%s", client.AWSClient.Region, task),
			})
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "ECS Task",
		"resources": len(resources),
	}).Debugf("Fetched resources")
	return resources, nil
}
