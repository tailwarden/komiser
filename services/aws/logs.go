package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

type LogsMetric struct {
	Metric     string
	Datapoints map[string]float64
}

func (aws AWS) MaximumLogsRetentionPeriod(cfg aws.Config) (int64, error) {
	var retention int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return retention, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := cloudwatchlogs.NewFromConfig(cfg)
		res, err := svc.DescribeLogGroups(context.Background(), &cloudwatchlogs.DescribeLogGroupsInput{})
		if err != nil {
			return retention, err
		}

		for _, group := range res.LogGroups {
			if group.RetentionInDays != nil && retention < int64(*group.RetentionInDays) {
				retention = int64(*group.RetentionInDays)
			}
		}
	}
	return retention, nil
}

func (awsClient AWS) GetLogsVolume(cfg aws.Config) ([]LogsMetric, error) {
	data := []LogsMetric{
		LogsMetric{
			Metric:     "IncomingBytes",
			Datapoints: map[string]float64{},
		},
		LogsMetric{
			Metric:     "IncomingLogEvents",
			Datapoints: map[string]float64{},
		},
	}

	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return data, err
	}

	for _, region := range regions {
		cfg.Region = region.Name
		cloudwatchClient := cloudwatch.NewFromConfig(cfg)
		resultCloudWatch, err := cloudwatchClient.GetMetricStatistics(context.Background(), &cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/Logs"),
			MetricName: aws.String("IncomingBytes"),
			StartTime:  aws.Time(time.Now().AddDate(0, 0, -7)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int32(86400),
			Statistics: []types.Statistic{
				types.StatisticSum,
			},
		})
		if err != nil {
			return data, err
		}

		for _, metric := range resultCloudWatch.Datapoints {
			key := (*metric.Timestamp).Format("2006-01-02")
			data[0].Datapoints[key] += *metric.Sum
		}

		resultCloudWatch2, err := cloudwatchClient.GetMetricStatistics(context.Background(), &cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/Logs"),
			MetricName: aws.String("IncomingLogEvents"),
			StartTime:  aws.Time(time.Now().AddDate(0, 0, -7)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int32(86400),
			Statistics: []types.Statistic{
				types.StatisticSum,
			},
		})
		if err != nil {
			return data, err
		}

		for _, metric := range resultCloudWatch2.Datapoints {
			key := (*metric.Timestamp).Format("2006-01-02")
			data[1].Datapoints[key] += *metric.Sum
		}
	}

	return data, nil
}
