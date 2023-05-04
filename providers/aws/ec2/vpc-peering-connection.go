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

func VpcPeeringConnections(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config ec2.DescribeVpcPeeringConnectionsInput
	resources := make([]Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	for {
		output, err := ec2Client.DescribeVpcPeeringConnections(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, vpcPeeringConnection := range output.VpcPeeringConnections {
			name := ""
			tags := make([]Tag, 0)
			for _, tag := range vpcPeeringConnection.Tags {
				if *tag.Key == "Name" {
					name = *tag.Value
				}
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "VPC Peering Connection",
				Region:     client.AWSClient.Region,
				Name:       name,
				ResourceId: *vpcPeeringConnection.VpcPeeringConnectionId,
				FetchedAt:  time.Now(),
				Cost:       0,
				Tags:       tags,
				Link: fmt.Sprintf(
					"https:/%s.console.aws.amazon.com/vpc/home?region=%s#PeeringConnectionDetails:VpcPeeringConnectionId=%s",
					client.AWSClient.Region,
					client.AWSClient.Region,
					*vpcPeeringConnection.VpcPeeringConnectionId,
				),
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
		"service":   "VPC Peering Connection",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
