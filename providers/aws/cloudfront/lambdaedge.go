package cloudfront

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	pricingTypes "github.com/aws/aws-sdk-go-v2/service/pricing/types"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

const (
	perOneMillonRequest = 1000000
)
func LambdaEdgeFunctions(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config cloudfront.ListFunctionsInput
	cloudfrontClient := cloudfront.NewFromConfig(*client.AWSClient)
	tempRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)
	pricingClient := pricing.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = tempRegion

	pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
		ServiceCode: aws.String("AmazonCloudFront"),
		Filters: []pricingTypes.Filter{
			{
				Field: aws.String("regionCode"),
				Value: aws.String(client.AWSClient.Region),
				Type:  pricingTypes.FilterTypeTermMatch,
			},
		},
	})
	if err != nil {
		log.Errorf("ERROR: Couldn't fetch pricing info for AWS CloudFront: %v", err)
	}

	priceMap, err := awsUtils.GetPriceMap(pricingOutput, "group")
	if err != nil {
		log.Errorf("ERROR: Failed to calculate cost per month: %v", err)
	}

	for {
		output, err := cloudfrontClient.ListFunctions(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, function := range output.FunctionList.Items {
			metricsLambdaEdgeDurationOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("Duration"),
				Namespace:  aws.String("AWS/CloudFront"),
				Dimensions: []types.Dimension{
					{
						Name:  aws.String("FunctionName"),
						Value: function.Name,
					},
				},
				Period: aws.Int32(86400),
				Statistics: []types.Statistic{
					types.StatisticAverage,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch Lambda@Edge Duration metric for %s", *function.Name)
			}

			lambdaEdgeDuration := 0.0
			if metricsLambdaEdgeDurationOutput != nil && len(metricsLambdaEdgeDurationOutput.Datapoints) > 0 {
				lambdaEdgeDuration = *metricsLambdaEdgeDurationOutput.Datapoints[0].Average
			}

			metricsLambdaEdgeRequestsOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("Requests"),
				Namespace:  aws.String("AWS/CloudFront"),
				Dimensions: []types.Dimension{
					{
						Name:  aws.String("FunctionName"),
						Value: function.Name,
					},
				},
				Period: aws.Int32(86400),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch Lambda@Edge Requests metric for %s", *function.Name)
			}

			lambdaEdgeRequests := 0.0
			if metricsLambdaEdgeRequestsOutput != nil && len(metricsLambdaEdgeRequestsOutput.Datapoints) > 0 {
				lambdaEdgeRequests = *metricsLambdaEdgeRequestsOutput.Datapoints[0].Sum
			}

			lambdaEdgeDurationCost := awsUtils.GetCost(priceMap["AWS-Lambda-Edge-Duration"], lambdaEdgeDuration)

			lambdaEdgeRequestsCost := awsUtils.GetCost(priceMap["AWS-Lambda-Edge-Requests"], lambdaEdgeRequests/perOneMillonRequest)

			monthlyCost := lambdaEdgeDurationCost + lambdaEdgeRequestsCost

			outputTags, err := cloudfrontClient.ListTagsForResource(ctx, &cloudfront.ListTagsForResourceInput{
				Resource: function.FunctionMetadata.FunctionARN,
			})

			tags := make([]Tag, 0)

			if err == nil {
				for _, tag := range outputTags.Tags.Items {
					tags = append(tags, Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "CloudFront",
				ResourceId: *function.FunctionMetadata.FunctionARN,
				Region:     client.AWSClient.Region,
				Name:       *function.Name,
				Cost:       monthlyCost,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/cloudfront/v3/home?region=%s#/distributions/%s", client.AWSClient.Region, client.AWSClient.Region, *function.Name),
			})
		}

		if aws.ToString(output.FunctionList.NextMarker) == "" {
			break
		}
		config.Marker = output.FunctionList.NextMarker
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "CloudFront",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}