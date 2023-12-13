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
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

func Distributions(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config cloudfront.ListDistributionsInput
	cloudfrontClient := cloudfront.NewFromConfig(*client.AWSClient)

	tempRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = tempRegion

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "Amazon CloudFront")
	if err != nil {
		log.Warnln("Couldn't fetch Amazon CloudFront cost and usage:", err)
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
					types.Dimension{
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

			metricsBytesUploadedOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("BytesUploaded"),
				Namespace:  aws.String("AWS/CloudFront"),
				Dimensions: []types.Dimension{
					types.Dimension{
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

			metricsRequestsOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("Requests"),
				Namespace:  aws.String("AWS/CloudFront"),
				Dimensions: []types.Dimension{
					types.Dimension{
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

			// calculate region data transfer out to internet
			dataTransferToInternet := (bytesUploaded / 1000000000) * 0.085

			// calculate region data transfer out to origin
			dataTransferToOrigin := (bytesDownloaded / 1000000000) * 0.02

			// calculate requests cost
			requestsCost := requests * 0.000001

			monthlyCost := dataTransferToInternet + dataTransferToOrigin + requestsCost

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
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
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
