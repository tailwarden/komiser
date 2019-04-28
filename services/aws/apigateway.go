package aws

import (
	"context"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) GetRestAPIs(cfg aws.Config) (int, error) {
	total := 0

	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}

	for _, region := range regions {
		cfg.Region = region.Name
		apigatewayClient := apigateway.New(cfg)
		req := apigatewayClient.GetRestApisRequest(&apigateway.GetRestApisInput{})
		res, err := req.Send(context.Background())
		if err != nil {
			return total, err
		}

		total += len(res.Items)
	}

	return total, nil
}

type APIGatewayMetric struct {
	Region     string
	Datapoints []Datapoint
}

func (awsClient AWS) GetAPIGatewayRequests(cfg aws.Config) ([]APIGatewayMetric, error) {
	metrics := make([]APIGatewayMetric, 0)

	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return metrics, err
	}

	for _, region := range regions {
		cfg.Region = region.Name
		cloudwatchClient := cloudwatch.New(cfg)
		reqCloudwatch := cloudwatchClient.GetMetricStatisticsRequest(&cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/ApiGateway"),
			MetricName: aws.String("Count"),
			StartTime:  aws.Time(time.Now().AddDate(0, -6, 0)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int64(86400),
			Statistics: []cloudwatch.Statistic{
				cloudwatch.StatisticSum,
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

		metrics = append(metrics, APIGatewayMetric{
			Region:     region.Name,
			Datapoints: series,
		})
	}

	return metrics, nil
}
