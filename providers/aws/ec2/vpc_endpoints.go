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

func VpcEndpoints(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config ec2.DescribeVpcEndpointsInput
	resources := make([]Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	for {
		output, err := ec2Client.DescribeVpcEndpoints(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, vpcEndpoint := range output.VpcEndpoints {
			name := ""
			tags := make([]Tag, 0)
			for _, tag := range vpcEndpoint.Tags {
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
				Service:    "VPC Endpoint",
				Region:     client.AWSClient.Region,
				Name:       name,
				ResourceId: *vpcEndpoint.VpcEndpointId,
				CreatedAt:  *vpcEndpoint.CreationTimestamp,
				FetchedAt:  time.Now(),
				Cost:       0,
				Tags:       tags,
				Link: fmt.Sprintf(
					"https:/%s.console.aws.amazon.com/vpc/home?region=%s#EndpointDetails:vpcEndpointId=%s",
					client.AWSClient.Region,
					client.AWSClient.Region,
					*vpcEndpoint.VpcEndpointId,
				),
				Metadata: map[string]string{
					"VpcEndpointType": string(vpcEndpoint.VpcEndpointType),
					"ServiceName":     string(*vpcEndpoint.ServiceName),
				},
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
		"service":   "VPC Endpoint",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
