package services

import (
	"fmt"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeS3Buckets(cfg aws.Config) (int, error) {
	svc := s3.New(cfg)
	req := svc.ListBucketsRequest(&s3.ListBucketsInput{})
	result, err := req.Send()
	if err != nil {
		return 0, err
	}
	return len(result.Buckets), nil
}

type BucketMetric struct {
	Bucket     string
	Datapoints []Datapoint
}

func (awsClient AWS) GetBucketsSize(cfg aws.Config) ([]BucketMetric, error) {
	metrics := make([]BucketMetric, 0)

	svc := s3.New(cfg)
	req := svc.ListBucketsRequest(&s3.ListBucketsInput{})
	result, err := req.Send()
	if err != nil {
		return metrics, err
	}

	for _, bucket := range result.Buckets {
		req := svc.GetBucketLocationRequest(&s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		result, err := req.Send()
		if err != nil {
			return metrics, err
		}

		fmt.Println(result)

		region, err := result.LocationConstraint.MarshalValue()
		if err != nil {
			return metrics, err
		}

		if len(region) > 0 {
			cfg.Region = region
		}

		fmt.Println(region)

		cloudwatchClient := cloudwatch.New(cfg)
		reqCloudwatch := cloudwatchClient.GetMetricStatisticsRequest(&cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/S3"),
			MetricName: aws.String("BucketSizeBytes"),
			StartTime:  aws.Time(time.Now().AddDate(0, -6, 0)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int64(86400),
			Statistics: []cloudwatch.Statistic{
				cloudwatch.StatisticAverage,
			},
			Dimensions: []cloudwatch.Dimension{
				cloudwatch.Dimension{
					Name:  aws.String("BucketName"),
					Value: bucket.Name,
				},
				cloudwatch.Dimension{
					Name:  aws.String("StorageType"),
					Value: aws.String("StandardStorage"),
				},
			},
		})
		resultCloudWatch, err := reqCloudwatch.Send()
		if err != nil {
			return metrics, err
		}

		series := make([]Datapoint, 0)

		for _, metric := range resultCloudWatch.Datapoints {
			series = append(series, Datapoint{
				Timestamp: *metric.Timestamp,
				Value:     *metric.Average,
			})
		}

		sort.Slice(series, func(i, j int) bool {
			return series[i].Timestamp.Sub(series[j].Timestamp) < 0
		})

		metrics = append(metrics, BucketMetric{
			Bucket:     *bucket.Name,
			Datapoints: series,
		})
	}
	return metrics, nil
}

func (awsClient AWS) GetBucketsObjects(cfg aws.Config) ([]BucketMetric, error) {
	metrics := make([]BucketMetric, 0)

	svc := s3.New(cfg)
	req := svc.ListBucketsRequest(&s3.ListBucketsInput{})
	result, err := req.Send()
	if err != nil {
		return metrics, err
	}

	for _, bucket := range result.Buckets {
		req := svc.GetBucketLocationRequest(&s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		result, err := req.Send()
		if err != nil {
			return metrics, err
		}

		fmt.Println(result)

		region, err := result.LocationConstraint.MarshalValue()
		if err != nil {
			return metrics, err
		}

		if len(region) > 0 {
			cfg.Region = region
		}

		fmt.Println(region)

		cloudwatchClient := cloudwatch.New(cfg)
		reqCloudwatch := cloudwatchClient.GetMetricStatisticsRequest(&cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/S3"),
			MetricName: aws.String("NumberOfObjects"),
			StartTime:  aws.Time(time.Now().AddDate(0, -6, 0)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int64(86400), //day
			Statistics: []cloudwatch.Statistic{
				cloudwatch.StatisticSum,
			},
			Dimensions: []cloudwatch.Dimension{
				cloudwatch.Dimension{
					Name:  aws.String("BucketName"),
					Value: bucket.Name,
				},
				cloudwatch.Dimension{
					Name:  aws.String("StorageType"),
					Value: aws.String("AllStorageTypes"),
				},
			},
		})
		resultCloudWatch, err := reqCloudwatch.Send()
		if err != nil {
			return metrics, err
		}

		series := make([]Datapoint, 0)

		for _, metric := range resultCloudWatch.Datapoints {
			series = append(series, Datapoint{
				Timestamp: *metric.Timestamp,
				Value:     *metric.Sum,
			})
		}

		sort.Slice(series, func(i, j int) bool {
			return series[i].Timestamp.Sub(series[j].Timestamp) < 0
		})

		metrics = append(metrics, BucketMetric{
			Bucket:     *bucket.Name,
			Datapoints: series,
		})
	}
	return metrics, nil
}
