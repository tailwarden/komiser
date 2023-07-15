package ec2

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func KeyPairs(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	ec2Client := ec2.NewFromConfig(*client.AWSClient)
	input := &ec2.DescribeKeyPairsInput{}

	output, err := ec2Client.DescribeKeyPairs(ctx, input)
	if err != nil {
		return resources, err
	}

	for _, keypair := range output.KeyPairs {
		tags := make([]models.Tag, 0)
		for _, tag := range keypair.Tags {
			tags = append(tags, models.Tag{
				Key:   aws.ToString(tag.Key),
				Value: aws.ToString(tag.Value),
			})
		}
			
		resources = append(resources, models.Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "EC2 KeyPair",
			ResourceId: aws.ToString(keypair.KeyPairId),
			Region:     client.AWSClient.Region,
			Name:       aws.ToString(keypair.KeyName),
			Cost:       0,
			Tags:       tags,
			CreatedAt:  *keypair.CreateTime,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/v2/home?region=%s#KeyPairs:search=%s", client.AWSClient.Region, client.AWSClient.Region, aws.ToString(keypair.KeyPairId)),
			Metadata: map[string]string{
				"KeyType": string(keypair.KeyType),
			},
		})
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"service":   "EC2 KeyPair",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
