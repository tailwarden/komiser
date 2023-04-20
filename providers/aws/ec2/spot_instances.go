package ec2

import (
	"context"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func SpotInstanceRequests(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	input := &ec2.DescribeSpotInstanceRequestsInput{}
	for {
		output, err := ec2Client.DescribeSpotInstanceRequests(ctx, input)
		if err != nil {
			return resources, err
		}

		for _, spotInstance := range output.SpotInstanceRequests {
			var tags []models.Tag
			for _, tag := range spotInstance.Tags {
				tags = append(tags, models.Tag{
					Key:   aws.ToString(tag.Key),
					Value: aws.ToString(tag.Value),
				})
			}

			cost := float64(0)
			if spotInstance.SpotPrice != nil {
				spotPrice, err := strconv.ParseFloat(*spotInstance.SpotPrice, 64)
				if err != nil {
					return resources, err
				}
				cost = spotPrice
			}

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "EC2 Spot Instance Request",
				ResourceId: aws.ToString(spotInstance.SpotInstanceRequestId),
				Region:     client.AWSClient.Region,
				Name:       aws.ToString(spotInstance.SpotInstanceRequestId),
				Cost:       cost,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#SpotInstancesDetails:id=%s", client.AWSClient.Region, client.AWSClient.Region, aws.ToString(spotInstance.SpotInstanceRequestId)),
				Metadata: map[string]string{
					"Availability Zone": aws.ToString(spotInstance.LaunchSpecification.Placement.AvailabilityZone),
				},
			})
		}

		if output.NextToken == nil {
			break
		}

		input.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "EC2 Spot Instance Request",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
