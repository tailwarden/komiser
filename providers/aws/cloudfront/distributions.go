package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
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
				Provider:  "AWS",
				Account:   client.Name,
				Service:   "CloudFront",
				Region:    client.AWSClient.Region,
				Name:      *distribution.DomainName,
				Cost:      0,
				Tags:      tags,
				FetchedAt: time.Now(),
			})
		}

		if aws.ToString(output.DistributionList.NextMarker) == "" {
			break
		}
		config.Marker = output.DistributionList.Marker
	}
	log.Printf("[%s] Fetched %d AWS Cloudfront distributions from %s\n", client.Name, len(resources), client.AWSClient.Region)
	return resources, nil
}
