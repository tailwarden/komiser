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
)

func Clusters(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecs.ListClustersInput
	ecsClient := ecs.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	for {
		output, err := ecsClient.ListClusters(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, cluster := range output.ClusterArns {
			resourceArn := fmt.Sprintf("arn:aws:ecs:%s:%s:cluster/%s", client.AWSClient.Region, *accountId, cluster)

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "ECS Cluster",
				ResourceId: resourceArn,
				Region:     client.AWSClient.Region,
				Name:       cluster,
				Cost:       0,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ecs/home?#/clusters/%s", client.AWSClient.Region, cluster),
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
		"service":   "ECS Cluster",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
