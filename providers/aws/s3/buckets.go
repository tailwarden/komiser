package s3

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func Buckets(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	if client.AWSClient.Region == "us-east-1" {
		var config s3.ListBucketsInput
		s3Client := s3.NewFromConfig(*client.AWSClient)
		output, err := s3Client.ListBuckets(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, bucket := range output.Buckets {
			tagsResp, err := s3Client.GetBucketTagging(context.Background(), &s3.GetBucketTaggingInput{
				Bucket: bucket.Name,
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

			resourceArn := fmt.Sprintf("arn:aws:s3:::%s", *bucket.Name)

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "S3",
				Region:     client.AWSClient.Region,
				ResourceId: resourceArn,
				Name:       *bucket.Name,
				Cost:       0,
				CreatedAt:  *bucket.CreationDate,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://s3.console.aws.amazon.com/s3/buckets/%s", *bucket.Name),
			})
		}
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "S3",
		"resources": len(resources),
	}).Debugf("Fetched resources")
	return resources, nil
}
