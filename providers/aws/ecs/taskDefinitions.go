package ecs

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	log "github.com/sirupsen/logrus"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func TaskDefinitions(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecs.ListTaskDefinitionsInput
	ecsClient := ecs.NewFromConfig(*client.AWSClient)
	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsoutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}
	accountId := stsoutput.Account

	for {
		output, err := ecsClient.ListTaskDefinitions(context.Background(), &config)
		if err != nil {
			return resources, err
		}
		for _, taskdefinition := range output.TaskDefinitionArns {
			resourceArn := fmt.Sprintf("arn:aws:ecs:%s:%s:taskdefinition/%s", client.AWSClient.Region, *accountId, taskdefinition)
			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "ECS Task Definition",
				ResourceId: resourceArn,
				Region:     client.AWSClient.Region,
				Name:       taskdefinition,
				Cost:       0,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ecs/home?#/taskdefinition/%s", client.AWSClient.Region, taskdefinition),
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
