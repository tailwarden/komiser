package aws

import (
	"context"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	models "github.com/mlabouardy/komiser/models/aws"
)

func (awsClient AWS) GetRestAPIs(cfg aws.Config) (int, error) {
	total := 0

	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return 0, err
	}

	for _, region := range regions {
		cfg.Region = region.Name
		apigatewayClient := apigateway.NewFromConfig(cfg)
		res, err := apigatewayClient.GetRestApis(context.Background(), &apigateway.GetRestApisInput{})
		if err != nil {
			return total, err
		}

		total += len(res.Items)
	}

	return total, nil
}

type APIGatewayMetric struct {
	Region     string
	Datapoints []models.Datapoint
}

func (awsClient AWS) GetAPIGatewayRequests(cfg aws.Config) ([]APIGatewayMetric, error) {
	metrics := make([]APIGatewayMetric, 0)

	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return metrics, err
	}

	for _, region := range regions {
		cfg.Region = region.Name
		cloudwatchClient := cloudwatch.NewFromConfig(cfg)
		resultCloudWatch, err := cloudwatchClient.GetMetricStatistics(context.Background(), &cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/ApiGateway"),
			MetricName: aws.String("Count"),
			StartTime:  aws.Time(time.Now().AddDate(0, -6, 0)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int32(86400),
			Statistics: []types.Statistic{
				types.StatisticSum,
			},
		})
		if err != nil {
			return metrics, err
		}

		series := make([]models.Datapoint, 0)

		for _, metric := range resultCloudWatch.Datapoints {
			series = append(series, models.Datapoint{
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
