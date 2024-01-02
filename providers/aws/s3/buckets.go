package s3

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	cloudwatchTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

func ConvertBytesToTerabytes(bytes int64) float64 {
	return float64(bytes) / 1099511627776
}

func Buckets(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config s3.ListBucketsInput
	s3Client := s3.NewFromConfig(*client.AWSClient)
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)

	tempRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	pricingClient := pricing.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = tempRegion
	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "S3")
	if err != nil {
		log.Warnln("Couldn't fetch S3 cost and usage:", err)
	}

	pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
		ServiceCode: aws.String("AmazonS3"),
		Filters: []types.Filter{
			{
				Field: aws.String("regionCode"),
				Value: aws.String(client.AWSClient.Region),
				Type:  types.FilterTypeTermMatch,
			},
		},
	})
	if err != nil {
		log.Errorf("ERROR: Couldn't fetch pricing info for AWS S3: %v", err)
		return resources, err
	}

	priceMap, err := awsUtils.GetPriceMap(pricingOutput, "group")
	if err != nil {
		log.Errorf("ERROR: Failed to calculate cost per month: %v", err)
		return resources, err
	}

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
			Dimensions: []cloudwatchTypes.Dimension{
				{
					Name:  aws.String("BucketName"),
					Value: bucket.Name,
				},
			},
			Unit:   cloudwatchTypes.StandardUnitBytes,
			Period: aws.Int32(3600),
			Statistics: []cloudwatchTypes.Statistic{
				cloudwatchTypes.StatisticAverage,
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

		metricsUsageOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
			StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
			EndTime:    aws.Time(time.Now()),
			MetricName: aws.String("AllRequests"),
			Namespace:  aws.String("AWS/S3"),
			Dimensions: []cloudwatchTypes.Dimension{
				{
					Name:  aws.String("BucketName"),
					Value: bucket.Name,
				},
			},
			Unit:   cloudwatchTypes.StandardUnitCount,
			Period: aws.Int32(3600),
			Statistics: []cloudwatchTypes.Statistic{
				cloudwatchTypes.StatisticSum,
			},
		})
		if err != nil {
			log.Warnf("Couldn't fetch usage metric for %s", *bucket.Name)
		}

		requestCount := 0.0
		if metricsUsageOutput != nil && len(metricsUsageOutput.Datapoints) > 0 {
			requestCount = *metricsUsageOutput.Datapoints[0].Sum
		}

		monthlyCost := 0.0

		storageCharges := awsUtils.GetCost(priceMap["AWS-S3-Storage"], sizeInTB*1024)
		requestCharges := awsUtils.GetCost(priceMap["S3-API-Tier1"], requestCount/1000)
		monthlyCost = storageCharges + requestCharges
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
			Metadata: map[string]string{
				"serviceCost": fmt.Sprint(serviceCost),
			},
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
