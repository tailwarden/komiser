package cloudfront

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func Distributions(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config cloudfront.ListDistributionsInput
	cloudfrontClient := cloudfront.NewFromConfig(*client.AWSClient)
	for {
		output, err := cloudfrontClient.ListDistributions(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, distribution := range output.DistributionList.Items {
			outputTags, err := cloudfrontClient.ListTagsForResource(ctx, &cloudfront.ListTagsForResourceInput{
				Resource: distribution.ARN,
			})

			tags := make([]Tag, 0)

			if err == nil {
				for _, tag := range outputTags.Tags.Items {
					tags = append(tags, Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "CloudFront",
				ResourceId: *distribution.ARN,
				Region:     client.AWSClient.Region,
				Name:       *distribution.DomainName,
				Cost:       0,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/cloudfront/v3/home?region=%s#/distributions/%s", client.AWSClient.Region, client.AWSClient.Region, *distribution.Id),
			})
		}

		if aws.ToString(output.DistributionList.NextMarker) == "" {
			break
		}
		config.Marker = output.DistributionList.Marker
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "CloudFront",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
