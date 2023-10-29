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
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils"
)

const (
	freeTierInvocations = 2000000
	costPerInvocation = 0.0000001
)

func Functions(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config cloudfront.ListFunctionsInput
	cloudfrontClient := cloudfront.NewFromConfig(*client.AWSClient)

	tempRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = tempRegion

	for {
		output, err := cloudfrontClient.ListFunctions(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, function := range output.FunctionList.Items {
			metricsInvocationsOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("FunctionInvocations"),
				Namespace:  aws.String("AWS/CloudFront"),
				Dimensions: []types.Dimension{
					types.Dimension{
						Name:  aws.String("FunctionName"),
						Value: function.Name,
					},
				},
				Period: aws.Int32(3600),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch invocations metric for %s", *function.Name)
				return resources, err
			}

			invocations := 0.0
			if metricsInvocationsOutput != nil && len(metricsInvocationsOutput.Datapoints) > 0 {
				invocations = *metricsInvocationsOutput.Datapoints[0].Sum
			}
			if invocations > freeTierInvocations {
				invocations -= freeTierInvocations
			}

			monthlyCost := invocations * costPerInvocation

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
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/cloudfront/v3/home?region=%s#/functions/%s", client.AWSClient.Region, client.AWSClient.Region, *function.Name),
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
