package cloudwatch

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	cwTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"

	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func getRate(pricingOutput *pricing.GetProductsOutput) (float64, error) {
	costPerMonth := 0.0

	if pricingOutput != nil && len(pricingOutput.PriceList) > 0 {
		var priceList interface{}
		err := json.Unmarshal([]byte(pricingOutput.PriceList[0]), &priceList)
		if err != nil {
			return 0, fmt.Errorf("failed to unmarshal JSON: %w", err)
		}

		priceListMap := priceList.(map[string]interface{})
		if onDemand, ok := priceListMap["terms"].(map[string]interface{})["OnDemand"]; ok {
			for _, details := range onDemand.(map[string]interface{}) {
				if priceDetails, ok := details.(map[string]interface{})["priceDimensions"].(map[string]interface{}); ok {
					for _, price := range priceDetails {
						usdPrice := price.(map[string]interface{})["pricePerUnit"].(map[string]interface{})["USD"].(string)
						costPerMonth, err = strconv.ParseFloat(usdPrice, 64)
						if err != nil {
							return 0, fmt.Errorf("failed to parse cost per month: %w", err)
						}
						break
					}
				}
			}
		}
	}

	return costPerMonth, nil
}

func MetricStreams(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	cloudWatchMetricsClient := cloudwatch.NewFromConfig(*client.AWSClient)

	tempRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	pricingClient := pricing.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = tempRegion

	pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
		ServiceCode: aws.String("AmazonCloudWatch"),
		Filters: []types.Filter{
			{
				Field: aws.String("regionCode"),
				Value: aws.String(client.AWSClient.Region),
				Type:  types.FilterTypeTermMatch,
			},
			{
				Field: aws.String("operation"),
				Value: aws.String("MetricUpdate"),
				Type:  types.FilterTypeTermMatch,
			},
		},
		MaxResults: aws.Int32(1),
	})
	if err != nil {
		log.Printf("ERROR: Couldn't fetch pricing info for Metric Streams: %v", err)
		return resources, err
	}

	costPerUpdate, err := getRate(pricingOutput)
	if err != nil {
		log.Printf("ERROR: Failed to calculate cost per month: %v", err)
		return resources, err
	}

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "AmazonCloudWatch")
	if err != nil {
		log.Warnln("Couldn't fetch AmazonCloudWatch cost and usage:", err)
	}

	input := &cloudwatch.ListMetricStreamsInput{}
	for {
		output, err := cloudWatchMetricsClient.ListMetricStreams(ctx, input)
		if err != nil {
			return resources, err
		}

		for _, stream := range output.Entries {
			tags := make([]models.Tag, 0)

			streamArn := aws.ToString(stream.Arn)
			tagInput := &cloudwatch.ListTagsForResourceInput{
				ResourceARN: &streamArn,
			}

			tagOutput, err := cloudWatchMetricsClient.ListTagsForResource(ctx, tagInput)
			if err == nil {
				for _, tag := range tagOutput.Tags {
					tags = append(tags, models.Tag{
						Key:   aws.ToString(tag.Key),
						Value: aws.ToString(tag.Value),
					})
				}
			}

			statisticsOutput, err := cloudWatchMetricsClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				MetricName: aws.String("MetricUpdate"),
				Namespace:  aws.String("AWS/CloudWatch/MetricStreams"),
				Dimensions: []cwTypes.Dimension{
					{
						Name:  aws.String("MetricStreamName"),
						Value: stream.Name,
					},
				},
			})
			if err != nil {
				return resources, err
			}

			updateCount := *statisticsOutput.Datapoints[0].Sum
			monthlyCost := costPerUpdate * updateCount

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "CloudWatch Metric Stream",
				ResourceId: streamArn,
				Region:     client.AWSClient.Region,
				Name:       aws.ToString(stream.Name),
				Cost:       monthlyCost,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				Link: fmt.Sprintf("https://%s.console.aws.amazon.com/cloudwatch/home?region=%s#metric-streams:streamsList/%s", client.AWSClient.Region, client.AWSClient.Region, aws.ToString(stream.Name)),
			})
		}

		if output.NextToken == nil {
			break
		}

		input.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "CloudWatch Metric Stream",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
