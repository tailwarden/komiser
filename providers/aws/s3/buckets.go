package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	. "github.com/mlabouardy/komiser/models"
)

func Buckets(ctx context.Context, cfg aws.Config, account string) ([]Resource, error) {
	resources := make([]Resource, 0)
	if cfg.Region != "us-east-1" {
		var config s3.ListBucketsInput
		s3Client := s3.NewFromConfig(cfg)
		output, err := s3Client.ListBuckets(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, o := range output.Buckets {
			tagsResp, err := s3Client.GetBucketTagging(context.Background(), &s3.GetBucketTaggingInput{
				Bucket: o.Name,
			})

			tags := make([]Tag, 0)
			if err == nil {
				for _, t := range tagsResp.TagSet {
					tags = append(tags, Tag{
						Key:   *t.Key,
						Value: *t.Value,
					})
				}
			}

			resources = append(resources, Resource{
				Provider:  "AWS",
				Account:   account,
				Service:   "S3",
				Region:    cfg.Region,
				Name:      *o.Name,
				Cost:      0,
				CreatedAt: *o.CreationDate,
				Tags:      tags,
				FetchedAt: time.Now(),
			})
		}
	}
	log.Printf("[%s] Fetched %d AWS S3 buckets from %s\n", account, len(resources), cfg.Region)
	return resources, nil
}
