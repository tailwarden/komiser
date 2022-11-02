package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	. "github.com/mlabouardy/komiser/models"
)

func Distributions(ctx context.Context, cfg aws.Config, account string) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config cloudfront.ListDistributionsInput
	cloudfrontClient := cloudfront.NewFromConfig(cfg)
	for {
		output, err := cloudfrontClient.ListDistributions(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, distribution := range output.DistributionList.Items {
			resources = append(resources, Resource{
				Provider:  "AWS",
				Account:   account,
				Service:   "CloudFront",
				Region:    cfg.Region,
				Name:      *distribution.DomainName,
				Cost:      0,
				FetchedAt: time.Now(),
			})

			if aws.ToString(output.DistributionList.NextMarker) == "" {
				break
			}
			config.Marker = output.DistributionList.Marker
		}
	}
	log.Printf("[%s] Fetched %d AWS Cloudfront distributions from %s\n", account, len(resources), cfg.Region)
	return resources, nil
}
