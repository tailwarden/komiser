package apigateway

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

func Apis(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config apigateway.GetRestApisInput
	apigatewayClient := apigateway.NewFromConfig(*client.AWSClient)
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)

	output, err := apigatewayClient.GetRestApis(ctx, &config)
	if err != nil {
		return resources, err
	}

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "Amazon API Gateway")
	if err != nil {
		log.Warnln("Couldn't fetch Amazon API Gateway cost and usage:", err)
	}

	for _, api := range output.Items {
		tags := make([]Tag, 0)
		for key, value := range api.Tags {
			tags = append(tags, Tag{
				Key:   key,
				Value: value,
			})
		}

		metricsCountOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
			StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
			EndTime:    aws.Time(time.Now()),
			MetricName: aws.String("Count"),
			Namespace:  aws.String("AWS/ApiGateway"),
			Dimensions: []types.Dimension{
				types.Dimension{
					Name:  aws.String("ApiName"),
					Value: api.Name,
				},
			},
			Period: aws.Int32(3600),
			Statistics: []types.Statistic{
				types.StatisticSum,
			},
		})

		if err != nil {
			log.Warnf("Couldn't fetch count metric for %s", *api.Name)
		}

		count := 0.0
		if metricsCountOutput != nil && len(metricsCountOutput.Datapoints) > 0 {
			count = *metricsCountOutput.Datapoints[0].Sum
		}

		monthlyCost := (count / 1000000)

		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "API Gateway",
			ResourceId: *api.Id,
			Region:     client.AWSClient.Region,
			Name:       *api.Name,
			Cost:       monthlyCost,
			Metadata: map[string]string{
				"serviceCost": fmt.Sprint(serviceCost),
			},
			Tags:       tags,
			CreatedAt:  *api.CreatedDate,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/apigateway/home?region=%s#/apis/%s", client.AWSClient.Region, client.AWSClient.Region, *api.Id),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "API Gateway",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
