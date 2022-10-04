package aws

import (
	"context"
	"fmt"
	"time"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	. "github.com/mlabouardy/komiser/models/aws"
)

type S3Bucket struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	Tags      []string  `json:"tags"`
	Region    string    `json:"region"`
}

func (aws AWS) DescribeS3Buckets(cfg awsConfig.Config) ([]S3Bucket, error) {
	buckets := make([]S3Bucket, 0)
	cfg.Region = "us-east-1"
	svc := s3.NewFromConfig(cfg)
	result, err := svc.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		return buckets, err
	}

	for _, bucket := range result.Buckets {
		bucketLocationResp, err := svc.GetBucketLocation(context.Background(), &s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			return buckets, err
		}

		cfg.Region = string(bucketLocationResp.LocationConstraint)
		svc := s3.NewFromConfig(cfg)
		tagsResp, err := svc.GetBucketTagging(context.Background(), &s3.GetBucketTaggingInput{
			Bucket: bucket.Name,
		})

		tags := make([]string, 0)
		if err == nil {
			for _, t := range tagsResp.TagSet {
				tags = append(tags, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
			}
		}

		buckets = append(buckets, S3Bucket{
			Name:      *bucket.Name,
			CreatedAt: *bucket.CreationDate,
			Tags:      tags,
			Region:    string(bucketLocationResp.LocationConstraint),
		})

	}
	return buckets, nil
}

type BucketMetric struct {
	Bucket     string
	Datapoints []Datapoint
}

func (awsClient AWS) GetBucketsSize(cfg awsConfig.Config) (map[string]map[string]float64, error) {
	metrics := make(map[string]map[string]float64, 0)
	cfg.Region = "us-east-1"
	svc := s3.NewFromConfig(cfg)

	result, err := svc.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		return metrics, err
	}

	for _, bucket := range result.Buckets {
		result, err := svc.GetBucketLocation(context.Background(), &s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		if err == nil {
			if len(string(result.LocationConstraint)) > 0 {
				cfg.Region = string(result.LocationConstraint)
			}

			cloudwatchClient := cloudwatch.NewFromConfig(cfg)
			resultCloudWatch, err := cloudwatchClient.GetMetricStatistics(context.Background(), &cloudwatch.GetMetricStatisticsInput{
				Namespace:  aws.String("AWS/S3"),
				MetricName: aws.String("BucketSizeBytes"),
				StartTime:  aws.Time(time.Now().AddDate(0, 0, -7)),
				EndTime:    aws.Time(time.Now()),
				Period:     aws.Int32(86400),
				Statistics: []types.Statistic{
					types.StatisticAverage,
				},
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
			})
			if err != nil {
				return metrics, err
			}

			for _, metric := range resultCloudWatch.Datapoints {
				if metrics[cfg.Region] == nil {
					metrics[cfg.Region] = make(map[string]float64, 0)
					metrics[cfg.Region][(*metric.Timestamp).Format("2006-01-02")] = *metric.Average
				} else {
					metrics[cfg.Region][(*metric.Timestamp).Format("2006-01-02")] += *metric.Average
				}
			}
		}
	}
	return metrics, nil
}

func (awsClient AWS) GetBucketsObjects(cfg awsConfig.Config) (map[string]map[string]float64, error) {
	metrics := make(map[string]map[string]float64, 0)
	cfg.Region = "us-east-1"
	svc := s3.NewFromConfig(cfg)

	result, err := svc.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		return metrics, err
	}

	for _, bucket := range result.Buckets {
		result, err := svc.GetBucketLocation(context.Background(), &s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		if err == nil {
			if len(string(result.LocationConstraint)) > 0 {
				cfg.Region = string(result.LocationConstraint)
			}

			cloudwatchClient := cloudwatch.NewFromConfig(cfg)
			resultCloudWatch, err := cloudwatchClient.GetMetricStatistics(context.Background(), &cloudwatch.GetMetricStatisticsInput{
				Namespace:  aws.String("AWS/S3"),
				MetricName: aws.String("NumberOfObjects"),
				StartTime:  aws.Time(time.Now().AddDate(0, 0, -7)),
				EndTime:    aws.Time(time.Now()),
				Period:     aws.Int32(86400), //day
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
				Dimensions: []types.Dimension{
					types.Dimension{
						Name:  aws.String("BucketName"),
						Value: bucket.Name,
					},
					types.Dimension{
						Name:  aws.String("StorageType"),
						Value: aws.String("AllStorageTypes"),
					},
				},
			})
			if err != nil {
				return metrics, err
			}

			for _, metric := range resultCloudWatch.Datapoints {
				if metrics[cfg.Region] == nil {
					metrics[cfg.Region] = make(map[string]float64, 0)
					metrics[cfg.Region][(*metric.Timestamp).Format("2006-01-02")] = *metric.Sum
				} else {
					metrics[cfg.Region][(*metric.Timestamp).Format("2006-01-02")] += *metric.Sum
				}
			}
		}

	}
	return metrics, nil
}

func (awsClient AWS) GetEmptyBuckets(cfg awsConfig.Config) (float64, error) {
	total := 0.0
	cfg.Region = "us-east-1"
	svc := s3.NewFromConfig(cfg)
	result, err := svc.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		return total, err
	}

	for _, bucket := range result.Buckets {
		result, err := svc.GetBucketLocation(context.Background(), &s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		if err == nil {
			if len(string(result.LocationConstraint)) > 0 {
				cfg.Region = string(result.LocationConstraint)
			}

			cloudwatchClient := cloudwatch.NewFromConfig(cfg)
			resultCloudWatch, err := cloudwatchClient.GetMetricStatistics(context.Background(), &cloudwatch.GetMetricStatisticsInput{
				Namespace:  aws.String("AWS/S3"),
				MetricName: aws.String("NumberOfObjects"),
				StartTime:  aws.Time(time.Now().AddDate(0, -6, 0)),
				EndTime:    aws.Time(time.Now()),
				Period:     aws.Int32(86400), //day
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
				Dimensions: []types.Dimension{
					types.Dimension{
						Name:  aws.String("BucketName"),
						Value: bucket.Name,
					},
					types.Dimension{
						Name:  aws.String("StorageType"),
						Value: aws.String("AllStorageTypes"),
					},
				},
			})
			if err != nil {
				return total, err
			}

			sum := 0.0

			for _, metric := range resultCloudWatch.Datapoints {
				sum += *metric.Sum
			}

			if sum == 0 {
				total++
			}
		}
	}
	return total, nil
}
