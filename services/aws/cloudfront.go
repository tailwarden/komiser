package aws

import (
	"context"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeCloudFrontDistributions(cfg aws.Config) (int, error) {
	svc := cloudfront.New(cfg)
	req := svc.ListDistributionsRequest(&cloudfront.ListDistributionsInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return 0, err
	}
	return len(result.DistributionList.Items), nil
}

type CloudFrontMetric struct {
	Distribution string
	Datapoints   []Datapoint
}

func (awsClient AWS) GetCloudFrontRequests(cfg aws.Config) ([]CloudFrontMetric, error) {
	metrics := make([]CloudFrontMetric, 0)

	svc := cloudfront.New(cfg)
	req := svc.ListDistributionsRequest(&cloudfront.ListDistributionsInput{})
	res, err := req.Send(context.Background())
	if err != nil {
		return metrics, err
	}

	cfg.Region = "us-east-1"
	cloudwatchClient := cloudwatch.New(cfg)

	for _, distribution := range res.DistributionList.Items {
		reqCloudwatch := cloudwatchClient.GetMetricStatisticsRequest(&cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/CloudFront"),
			MetricName: aws.String("Requests"),
			StartTime:  aws.Time(time.Now().AddDate(0, -6, 0)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int64(86400),
			Statistics: []cloudwatch.Statistic{
				cloudwatch.StatisticSum,
			},
			Dimensions: []cloudwatch.Dimension{
				cloudwatch.Dimension{
					Name:  aws.String("DistributionId"),
					Value: distribution.Id,
				},
				cloudwatch.Dimension{
					Name:  aws.String("Region"),
					Value: aws.String("Global"),
				},
			},
		})

		resultCloudWatch, err := reqCloudwatch.Send(context.Background())
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
