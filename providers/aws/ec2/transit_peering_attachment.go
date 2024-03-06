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

func TransitGatewayPreeringAttachments(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config ec2.DescribeTransitGatewayPeeringAttachmentsInput
	resources := make([]Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	for {
		output, err := ec2Client.DescribeTransitGatewayPeeringAttachments(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, tAttachment := range output.TransitGatewayPeeringAttachments {
			tags := make([]Tag, 0)
			for _, tag := range tAttachment.Tags {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Transit Gateway Peering Attachments",
				Region:     client.AWSClient.Region,
				ResourceId: *tAttachment.TransitGatewayAttachmentId,
				Cost:       0,
				Name:       *tAttachment.TransitGatewayAttachmentId,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/vpc/home?region=%s#InternetGateway:internetGatewayId=%s", client.AWSClient.Region, client.AWSClient.Region),
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
		"service":   "Transit Gateway Peering Attachments",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
