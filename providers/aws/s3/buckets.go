package s3

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils"
)

func ConvertBytesToTerabytes(bytes int64) float64 {
	return float64(bytes) / 1000000000000
}

func Buckets(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config s3.ListBucketsInput
	s3Client := s3.NewFromConfig(*client.AWSClient)
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)
	output, err := s3Client.ListBuckets(context.Background(), &config)
	if err != nil {
		return resources, err
	}

	for _, bucket := range output.Buckets {
		metricsBucketSizebytesOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
			StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
			EndTime:    aws.Time(time.Now()),
			MetricName: aws.String("BucketSizeBytes"),
			Namespace:  aws.String("AWS/S3"),
			Dimensions: []types.Dimension{
				types.Dimension{
					Name:  aws.String("BucketName"),
					Value: bucket.Name,
				},
				types.Dimension{
					Name:  aws.String("StorageType"),
					Value: aws.String("StandardStorage"),
				},
			},
			Unit:   types.StandardUnitBytes,
			Period: aws.Int32(3600),
			Statistics: []types.Statistic{
				types.StatisticAverage,
			},
		})
		if err != nil {
			log.Warnf("Couldn't fetch invocations metric for %s", *bucket.Name)
		}

		bucketSize := 0.0
		if metricsBucketSizebytesOutput != nil && len(metricsBucketSizebytesOutput.Datapoints) > 0 {
			bucketSize = *metricsBucketSizebytesOutput.Datapoints[0].Average
		}

		sizeInTB := ConvertBytesToTerabytes(int64(bucketSize))
		monthlyCost := 0.0

		if sizeInTB <= 50 {
			monthlyCost = (sizeInTB * 1000) * 0.023
		} else if sizeInTB <= 450 {
			monthlyCost = (sizeInTB * 1000) * 0.022
		} else {
			monthlyCost = (sizeInTB * 1000) * 0.021
		}

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
			Cost:       monthlyCost,
			CreatedAt:  *bucket.CreationDate,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://s3.console.aws.amazon.com/s3/buckets/%s", *bucket.Name),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "S3",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
