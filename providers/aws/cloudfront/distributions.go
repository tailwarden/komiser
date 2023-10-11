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
	freeTierRequests           = 10000000
	freeTierFunctionInvocation = 2000000
	freeTierUpload             = 1099511627776
)

func ConvertBytesToTerabytes(bytes int64) float64 {
	return float64(bytes) / 1099511627776
}

func Distributions(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config cloudfront.ListDistributionsInput
	cloudfrontClient := cloudfront.NewFromConfig(*client.AWSClient)

	tempRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = tempRegion
	pricingClient := pricing.NewFromConfig(*client.AWSClient)

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
		return resources, err
	}

	priceMap, err := awsUtils.GetPriceMap(pricingOutput)
	if err != nil {
		log.Errorf("ERROR: Failed to calculate cost per month: %v", err)
		return resources, err
	}

	for {
		output, err := cloudfrontClient.ListDistributions(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, distribution := range output.DistributionList.Items {
			metricsBytesDownloadedOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("BytesDownloaded"),
				Namespace:  aws.String("AWS/CloudFront"),
				Dimensions: []types.Dimension{
					{
						Name:  aws.String("DistributionId"),
						Value: distribution.Id,
					},
				},
				Period: aws.Int32(86400),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch invocations metric for %s", *distribution.Id)
			}

			bytesDownloaded := 0.0
			if metricsBytesDownloadedOutput != nil && len(metricsBytesDownloadedOutput.Datapoints) > 0 {
				bytesDownloaded = *metricsBytesDownloadedOutput.Datapoints[0].Sum
			}

			sizeInTBDownload := ConvertBytesToTerabytes(int64(bytesDownloaded))

			metricsBytesUploadedOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("BytesUploaded"),
				Namespace:  aws.String("AWS/CloudFront"),
				Dimensions: []types.Dimension{
					{
						Name:  aws.String("DistributionId"),
						Value: distribution.Id,
					},
				},
				Period: aws.Int32(86400),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch invocations metric for %s", *distribution.Id)
			}

			bytesUploaded := 0.0
			if metricsBytesUploadedOutput != nil && len(metricsBytesUploadedOutput.Datapoints) > 0 {
				bytesUploaded = *metricsBytesUploadedOutput.Datapoints[0].Sum
			}
			if bytesUploaded > freeTierUpload {
				bytesUploaded -= freeTierUpload
			}

			sizeInTBUpload := ConvertBytesToTerabytes(int64(bytesUploaded))

			metricsRequestsOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("Requests"),
				Namespace:  aws.String("AWS/CloudFront"),
				Dimensions: []types.Dimension{
					{
						Name:  aws.String("DistributionId"),
						Value: distribution.Id,
					},
				},
				Period: aws.Int32(86400),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch invocations metric for %s", *distribution.Id)
			}

			requests := 0.0
			if metricsRequestsOutput != nil && len(metricsRequestsOutput.Datapoints) > 0 {
				requests = *metricsRequestsOutput.Datapoints[0].Sum
			}
			if requests > freeTierRequests {
				requests -= freeTierRequests
			}

			metricsLambdaEdgeDurationOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("Duration"),
				Namespace:  aws.String("AWS/LambdaEdge"),
				Dimensions: []types.Dimension{
					{
						Name:  aws.String("DistributionId"),
						Value: distribution.Id,
					},
				},
				Period: aws.Int32(86400),
				Statistics: []types.Statistic{
					types.StatisticAverage,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch Lambda@Edge Duration metric for %s", *distribution.Id)
			}

			lambdaEdgeDuration := 0.0
			if metricsLambdaEdgeDurationOutput != nil && len(metricsLambdaEdgeDurationOutput.Datapoints) > 0 {
				lambdaEdgeDuration = *metricsLambdaEdgeDurationOutput.Datapoints[0].Average
			}

			metricsLambdaEdgeRequestsOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("Requests"),
				Namespace:  aws.String("AWS/LambdaEdge"),
				Dimensions: []types.Dimension{
					{
						Name:  aws.String("DistributionId"),
						Value: distribution.Id,
					},
				},
				Period: aws.Int32(86400),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch Lambda@Edge Requests metric for %s", *distribution.Id)
			}

			lambdaEdgeRequests := 0.0
			if metricsLambdaEdgeRequestsOutput != nil && len(metricsLambdaEdgeRequestsOutput.Datapoints) > 0 {
				lambdaEdgeRequests = *metricsLambdaEdgeRequestsOutput.Datapoints[0].Sum
			}

			metricsFunctionInvocationsOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("FunctionInvocations"),
				Namespace:  aws.String("AWS/CloudFront"),
				Dimensions: []types.Dimension{
					{
						Name:  aws.String("DistributionId"),
						Value: distribution.Id,
					},
				},
				Period: aws.Int32(3600),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
			})
			if err != nil {
				log.Warnf("Couldn't fetch Function Invocations metric for %s", *distribution.Id)
			}

			functionInvocation := 0.0
			if metricsFunctionInvocationsOutput != nil && len(metricsFunctionInvocationsOutput.Datapoints) > 0 {
				functionInvocation = *metricsFunctionInvocationsOutput.Datapoints[0].Sum
			}
			if functionInvocation > freeTierFunctionInvocation {
				functionInvocation -= freeTierFunctionInvocation
			}

			dataTransferToInternetCost := awsUtils.GetCost(priceMap["CloudFront-DataTransfer-In-Bytes"], sizeInTBUpload*1024)

			dataTransferToOriginCost := awsUtils.GetCost(priceMap["CloudFront-DataTransfer-Out-Bytes"], sizeInTBDownload*1024)

			requestsCost := awsUtils.GetCost(priceMap["CloudFront-Requests"], requests/10000)

			lambdaEdgeDurationCost := awsUtils.GetCost(priceMap["AWS-Lambda-Edge-Duration"], lambdaEdgeDuration)

			lambdaEdgeRequestsCost := awsUtils.GetCost(priceMap["AWS-Lambda-Edge-Requests"], lambdaEdgeRequests/10000000)

			functionInvocationsCost := awsUtils.GetCost(priceMap["AWS-CloudFront-FunctionInvocation"], functionInvocation)

			monthlyCost := dataTransferToInternetCost + dataTransferToOriginCost + requestsCost + lambdaEdgeDurationCost + lambdaEdgeRequestsCost + functionInvocationsCost

			outputTags, err := cloudfrontClient.ListTagsForResource(ctx, &cloudfront.ListTagsForResourceInput{
				Resource: distribution.ARN,
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
				ResourceId: *distribution.ARN,
				Region:     client.AWSClient.Region,
				Name:       *distribution.DomainName,
				Cost:       monthlyCost,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/cloudfront/v3/home?region=%s#/distributions/%s", client.AWSClient.Region, client.AWSClient.Region, *distribution.Id),
			})
		}

		if aws.ToString(output.DistributionList.NextMarker) == "" {
			break
		}
		config.Marker = output.DistributionList.Marker
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
