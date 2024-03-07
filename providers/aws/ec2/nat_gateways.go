package ec2

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func NatGateways(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config ec2.DescribeNatGatewaysInput
	resources := make([]models.Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	for {
		output, err := ec2Client.DescribeNatGateways(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, natGateways := range output.NatGateways {
			tags := make([]models.Tag, 0)
			for _, tag := range natGateways.Tags {
				tags = append(tags, models.Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Nat Gateway",
				Region:     client.AWSClient.Region,
				ResourceId: *natGateways.NatGatewayId,
				Cost:       0,
				Name:       *natGateways.NatGatewayId,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/vpc/home?region=%s#NatGateway:natGatewayId=%s", client.AWSClient.Region, client.AWSClient.Region, *natGateways.NatGatewayId),
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
		"service":   "Nat Gateway",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
