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

func Services(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecs.ListServicesInput
	ecsClient := ecs.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	for {
		output, err := ecsClient.ListServices(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, service := range output.ServiceArns {
			resourceArn := fmt.Sprintf("arn:aws:ecs:%s:%s:service/%s", client.AWSClient.Region, *accountId, service)

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "ECS Service",
				ResourceId: resourceArn,
				Region:     client.AWSClient.Region,
				Name:       service,
				Cost:       0,
				FetchedAt:  time.Now(),
			})

		}
		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.Debugf("[%s] Fetched %d AWS ECS services from %s\n", client.Name, len(resources), client.AWSClient.Region)
	return resources, nil
}
