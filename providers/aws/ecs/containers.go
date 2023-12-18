package ecs

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func ContainerInstances(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecs.ListContainerInstancesInput
	ecsContainer := ecs.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})

	if err != nil {
		return resources, err
	}

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "ECS")
	if err != nil {
		log.Warnln("Couldn't fetch ECS cost and usage:", err)
	}
	accountId := stsOutput.Account
	for {
		output, err := ecsContainer.ListContainerInstances(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, containerInstance := range output.ContainerInstanceArns {
			resourceArn := fmt.Sprintf("arn:aws:ecs:%s:%s:instance/%s", client.AWSClient.Region, *accountId, containerInstance)

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "ECS Container Instance",
				ResourceId: resourceArn,
				Region:     client.AWSClient.Region,
				Name:       containerInstance,
				Cost:       0,
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ecs/home?#/containers/%s", client.AWSClient.Region, containerInstance),
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
		"service":   "ECS Container Instance",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
