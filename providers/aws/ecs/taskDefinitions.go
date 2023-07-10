package ecs

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	log "github.com/sirupsen/logrus"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func extractNameAndRevisionFromArn(input string) []string {
	var nameAndRevision [2]string

	parts := strings.Split(input, ":")
	if len(parts) >= 6 {
		nameAndRevision[0] = strings.Split(parts[5], "/")[1]
		nameAndRevision[1] = parts[6]
	}
	return nameAndRevision[:]
}

func TaskDefinitions(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecs.ListTaskDefinitionsInput
	ecsClient := ecs.NewFromConfig(*client.AWSClient)

	for {
		output, err := ecsClient.ListTaskDefinitions(context.Background(), &config)
		if err != nil {
			return resources, err
		}
		for _, taskdefinition := range output.TaskDefinitionArns {
			ecsNameAndRevision := extractNameAndRevisionFromArn(taskdefinition)
			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "ECS Task Definition",
				ResourceId: taskdefinition,
				Region:     client.AWSClient.Region,
				Name:       ecsNameAndRevision[0],
				Cost:       0,
				FetchedAt:  time.Now(),

				Link: fmt.Sprintf("https://%s.console.aws.amazon.com/ecs/v2/task-definitions/%s/%s/containers?region=%s", client.AWSClient.Region, ecsNameAndRevision[0], ecsNameAndRevision[1], client.AWSClient.Region),
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
		"service":   "ECS Task Definition",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
