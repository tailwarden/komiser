package ec2

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func InternetGateways(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config ec2.DescribeInternetGatewaysInput
	resources := make([]Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	for {
		output, err := ec2Client.DescribeInternetGateways(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, internetGateways := range output.InternetGateways {
			tags := make([]Tag, 0)
			for _, tag := range internetGateways.Tags {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Internet Gateway",
				Region:     client.AWSClient.Region,
				ResourceId: *internetGateways.InternetGatewayId,
				Cost:       0,
				Name:       *internetGateways.InternetGatewayId,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/vpc/home?region=%s#InternetGateway:internetGatewayId=%s", client.AWSClient.Region, client.AWSClient.Region, *internetGateways.InternetGatewayId),
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
		"service":   "Internet Gateway",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
