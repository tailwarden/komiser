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
		output, err := cloudfrontClient.ListDistributions(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, distribution := range output.DistributionList.Items {
			resources = append(resources, Resource{
				Provider:  "AWS",
				Account:   client.Name,
				Service:   "CloudFront",
				Region:    client.AWSClient.Region,
				Name:      *distribution.DomainName,
				Cost:      0,
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
