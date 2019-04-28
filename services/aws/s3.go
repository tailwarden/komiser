package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeS3Buckets(cfg aws.Config) (int, error) {
	svc := s3.New(cfg)
	req := svc.ListBucketsRequest(&s3.ListBucketsInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return 0, err
	}
	return len(result.Buckets), nil
}

type BucketMetric struct {
	Bucket     string
	Datapoints []Datapoint
}

func (awsClient AWS) GetBucketsSize(cfg aws.Config) (map[string]map[string]float64, error) {
	metrics := make(map[string]map[string]float64, 0)

	svc := s3.New(cfg)
	req := svc.ListBucketsRequest(&s3.ListBucketsInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return metrics, err
	}

	for _, bucket := range result.Buckets {
		req := svc.GetBucketLocationRequest(&s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		result, err := req.Send(context.Background())
		if err != nil {
			return metrics, err
		}

		region, err := result.LocationConstraint.MarshalValue()
		if err != nil {
			return metrics, err
		}

		if len(region) > 0 {
			cfg.Region = region
		}

		cloudwatchClient := cloudwatch.New(cfg)
		reqCloudwatch := cloudwatchClient.GetMetricStatisticsRequest(&cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/S3"),
			MetricName: aws.String("BucketSizeBytes"),
			StartTime:  aws.Time(time.Now().AddDate(0, 0, -7)),
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
		resultCloudWatch, err := reqCloudwatch.Send(context.Background())
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
	return metrics, nil
}

func (awsClient AWS) GetBucketsObjects(cfg aws.Config) (map[string]map[string]float64, error) {
	metrics := make(map[string]map[string]float64, 0)

	svc := s3.New(cfg)
	req := svc.ListBucketsRequest(&s3.ListBucketsInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return metrics, err
	}

	for _, bucket := range result.Buckets {
		req := svc.GetBucketLocationRequest(&s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		result, err := req.Send(context.Background())
		if err != nil {
			return metrics, err
		}

		region, err := result.LocationConstraint.MarshalValue()
		if err != nil {
			return metrics, err
		}

		if len(region) > 0 {
			cfg.Region = region
		}

		cloudwatchClient := cloudwatch.New(cfg)
		reqCloudwatch := cloudwatchClient.GetMetricStatisticsRequest(&cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/S3"),
			MetricName: aws.String("NumberOfObjects"),
			StartTime:  aws.Time(time.Now().AddDate(0, 0, -7)),
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
		resultCloudWatch, err := reqCloudwatch.Send(context.Background())
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
	return metrics, nil
}

func (awsClient AWS) GetEmptyBuckets(cfg aws.Config) (float64, error) {
	total := 0.0

	svc := s3.New(cfg)
	req := svc.ListBucketsRequest(&s3.ListBucketsInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return total, err
	}

	for _, bucket := range result.Buckets {
		req := svc.GetBucketLocationRequest(&s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		result, err := req.Send(context.Background())
		if err != nil {
			return total, err
		}

		region, err := result.LocationConstraint.MarshalValue()
		if err != nil {
			return total, err
		}

		if len(region) > 0 {
			cfg.Region = region
		}

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
		resultCloudWatch, err := reqCloudwatch.Send(context.Background())
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
	return total, nil
}
