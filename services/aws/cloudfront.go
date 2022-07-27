package aws

import (
	"context"
	"sort"
	"time"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go/aws"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeCloudFrontDistributions(cfg awsConfig.Config) (int, error) {
	svc := cloudfront.NewFromConfig(cfg)
	result, err := svc.ListDistributions(context.Background(), &cloudfront.ListDistributionsInput{})
	if err != nil {
		return 0, err
	}
	return len(result.DistributionList.Items), nil
}

type CloudFrontMetric struct {
	Distribution string
	Datapoints   []Datapoint
}

func (awsClient AWS) GetCloudFrontRequests(cfg awsConfig.Config) ([]CloudFrontMetric, error) {
	metrics := make([]CloudFrontMetric, 0)

	svc := cloudfront.NewFromConfig(cfg)
	res, err := svc.ListDistributions(context.Background(), &cloudfront.ListDistributionsInput{})
	if err != nil {
		return metrics, err
	}

	cfg.Region = "us-east-1"
	cloudwatchClient := cloudwatch.NewFromConfig(cfg)

	for _, distribution := range res.DistributionList.Items {
		resultCloudWatch, err := cloudwatchClient.GetMetricStatistics(context.Background(), &cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/CloudFront"),
			MetricName: aws.String("Requests"),
			StartTime:  aws.Time(time.Now().AddDate(0, -6, 0)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int32(86400),
			Statistics: []types.Statistic{
				types.StatisticSum,
			},
			Dimensions: []types.Dimension{
				types.Dimension{
					Name:  aws.String("DistributionId"),
					Value: distribution.Id,
				},
				types.Dimension{
					Name:  aws.String("Region"),
					Value: aws.String("Global"),
				},
			},
		})

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

		metrics = append(metrics, CloudFrontMetric{
			Distribution: *distribution.Id,
			Datapoints:   series,
		})
	}

	return metrics, nil
}
