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

func TransitGatewayVpcAttachments(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config ec2.DescribeTransitGatewayVpcAttachmentsInput
	resources := make([]models.Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	for {
		output, err := ec2Client.DescribeTransitGatewayVpcAttachments(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, tAttachment := range output.TransitGatewayVpcAttachments {
			tags := make([]models.Tag, 0)
			for _, tag := range tAttachment.Tags {
				tags = append(tags, models.Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Transit Gateway Vpc Attachments",
				Region:     client.AWSClient.Region,
				ResourceId: *tAttachment.TransitGatewayAttachmentId,
				Cost:       0,
				Name:       *tAttachment.TransitGatewayAttachmentId,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/vpc/home?region=%s#TransitGatewayAttachment:transitGatewayAttachmentId=%s", client.AWSClient.Region, client.AWSClient.Region, *tAttachment.TransitGatewayAttachmentId),
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
		"service":   "Transit Gateway Vpc Attachments",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
