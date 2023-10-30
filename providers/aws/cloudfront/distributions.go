package cloudfront

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/pricing"

	// pricingTypes "github.com/aws/aws-sdk-go-v2/service/pricing/types"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

const (
	freeTierRequests = 10000000
	freeTierUpload   = 1099511627776
	per10kRequest    = 10000
)

var EdgeLocation string

func Distributions(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config cloudfront.ListDistributionsInput
	cloudfrontClient := cloudfront.NewFromConfig(*client.AWSClient)
	pricingClient := pricing.NewFromConfig(*client.AWSClient)

	pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
		ServiceCode: aws.String("AmazonCloudFront"),
	})
	if err != nil {
		log.Errorf("ERROR: Couldn't fetch pricing info for AWS CloudFront: %v", err)
	}

	priceMapForDataTransfer, err := GetPriceMapCF(pricingOutput, "fromLocation")
	if err != nil {
		log.Errorf("ERROR: Failed to calculate cost per month: %v", err)
	}

	priceMapForRequest, err := GetPriceMapCF(pricingOutput, "location")
	if err != nil {
		log.Errorf("ERROR: Failed to calculate cost per month: %v", err)
	}

	for {
		getRegions := getRegionMapping()
		for region, locations := range getRegions {
			client.AWSClient.Region = region
			for _, edgelocation := range locations {
				output, err := cloudfrontClient.ListDistributions(ctx, &config)
				if err != nil {
					return resources, err
				}

				cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)

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

					if priceMapForDataTransfer[edgelocation] != nil {
						EdgeLocation = edgelocation
						break
					}
					dataTransferToOriginCost := awsUtils.GetCost(priceMapForDataTransfer[EdgeLocation], (float64(bytesDownloaded)/1099511627776)*1024)

					requestsCost := awsUtils.GetCost(priceMapForRequest[EdgeLocation], requests/per10kRequest)

					monthlyCost := dataTransferToOriginCost + requestsCost

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

	}
}

func getRegionMapping() map[string][]string {
	return map[string][]string{
		"us-east-1":      {"United States", "Mexico", "Canada"},
		"eu-west-1":      {"Europe", "Israel"},
		"ap-northeast-1": {"Australia", "New Zealand", "Taiwan"},
		"ap-northeast-2": {"South Korea"},
		"ap-southeast-1": {"Philippines", "Singapore", "Thailand", "Malaysia"},
		"ap-southeast-3": {"Indonesia"},
		"ap-south-1":     {"India"},
		"sa-east-1":      {"Japan", "South America"},
		"me-south-1":     {"South Africa", "Kenya", "Middle East"},
		"ap-east-1":      {"Hong Kong", "Vietnam"},
		"cn-north-1":     {"China"},
	}
}

// GetPriceMapCF is modified functions from awsUtils.GetPriceMap to get CF distribution unit price based on location
func GetPriceMapCF(pricingOutput *pricing.GetProductsOutput, field string) (map[string][]awsUtils.PriceDimensions, error) {
	priceMap := make(map[string][]awsUtils.PriceDimensions)

	if pricingOutput != nil && len(pricingOutput.PriceList) > 0 {
		for _, item := range pricingOutput.PriceList {
			price := awsUtils.ProductEntry{}
			err := json.Unmarshal([]byte(item), &price)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
			}

			var key string
			switch field {
			case "fromLocation":
				if price.Product.Attributes.TransferType == "CloudFront to Origin" {

					key = price.Product.Attributes.FromLocation
				}
			case "location":
				if price.Product.Attributes.RequestType == "CloudFront-Request-HTTP-Proxy" {
					key = price.Product.Attributes.RequestLocation
				}
			}

			unitPrices := []awsUtils.PriceDimensions{}
			for _, pd := range price.Terms.OnDemand {
				for _, p := range pd.PriceDimensions {
					unitPrices = append(unitPrices, p)
				}
			}

			priceMap[key] = unitPrices
		}
	}

	return priceMap, nil
}
