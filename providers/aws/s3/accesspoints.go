package s3

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func AccessPoints(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	s3ControlClient := s3control.NewFromConfig(cfg)
	stsClient := sts.NewFromConfig(cfg)
	identity, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}
	accountId := *identity.Account

	output, err := s3ControlClient.ListAccessPoints(ctx, &s3control.ListAccessPointsInput{
		AccountId: &accountId,
	})
	if err != nil {
		return resources, err
	}

	for _, accesspoints := range output.AccessPointList {
		objects, err := s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
			Bucket: accesspoints.Bucket,
		})

		if err != nil {
			log.Warnf("Couldn't fetch objects for %s", *accesspoints.Name)
			continue
		}

		var lastModified time.Time
		for _, object := range objects.Contents {
			if object.LastModified.After(lastModified) {
				lastModified = *object.LastModified
			}
		}
		tagsResp, err := s3Client.GetBucketTagging(ctx, &s3.GetBucketTaggingInput{
			Bucket: accesspoints.Bucket,
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

		resourceArn := fmt.Sprintf("arn:aws:s3:::%s", *accesspoints.Name)
		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "Accesspoints",
			Region:     client.AWSClient.Region,
			ResourceId: resourceArn,
			Name:       *accesspoints.Name,
			Cost:       0,
			CreatedAt:  lastModified,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://s3.console.aws.amazon.com/s3/buckets/%s", *accesspoints.Name),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Accesspoints",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
