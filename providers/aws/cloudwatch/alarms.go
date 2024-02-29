package cloudwatch

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

const (
	AverageHoursPerMonth = 730
)

func Alarms(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	var config cloudwatch.DescribeAlarmsInput
	// This code temporarily changes the region to "us-east-1" and creates a new Pricing client
	// then changes the region back to what it was before.
	// This is necessary because the Pricing client needs to operate in the "us-east-1" region
	oldRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	pricingClient := pricing.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = oldRegion
	cloudWatchClient := cloudwatch.NewFromConfig(*client.AWSClient)

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "AmazonCloudWatch")
	if err != nil {
		log.Warnln("Couldn't fetch AmazonCloudWatch cost and usage:", err)
	}
	for {
		output, err := cloudWatchClient.DescribeAlarms(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, alarm := range output.MetricAlarms {
			outputTags, err := cloudWatchClient.ListTagsForResource(ctx, &cloudwatch.ListTagsForResourceInput{
				ResourceARN: alarm.AlarmArn,
			})

			tags := make([]models.Tag, 0)

			if err == nil {
				for _, tag := range outputTags.Tags {
					tags = append(tags, models.Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}
			}

			pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
				ServiceCode: aws.String("AmazonCloudWatch"),
				Filters: []types.Filter{
					{
						Field: aws.String("regionCode"),
						Value: aws.String(client.AWSClient.Region),
						Type:  types.FilterTypeTermMatch,
					},
					{
						Field: aws.String("group"),
						Value: aws.String("Alarm"),
						Type:  types.FilterTypeTermMatch,
					},
				},
				MaxResults: aws.Int32(1),
			})
			if err != nil {
				log.Printf("ERROR: Couldn't fetch pricing info for alarm %s: %v", *alarm.AlarmName, err)
				continue
			}

			costPerMonth, err := calculateCostPerMonth(pricingOutput)
			if err != nil {
				log.Printf("ERROR: Failed to calculate cost per month: %v", err)
				continue
			}

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "CloudWatch",
				ResourceId: *alarm.AlarmArn,
				Region:     client.AWSClient.Region,
				Name:       *alarm.AlarmName,
				Cost:       costPerMonth,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				Link: fmt.Sprintf("https://%s.console.aws.amazon.com/cloudwatch/home?region=%s#alarmsV2:alarm/%s", client.AWSClient.Region, client.AWSClient.Region, *alarm.AlarmName),
			})
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "CloudWatch",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}

func calculateCostPerMonth(pricingOutput *pricing.GetProductsOutput) (float64, error) {
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
						cost, err := strconv.ParseFloat(usdPrice, 64)
						if err != nil {
							return 0, fmt.Errorf("failed to parse cost per month: %w", err)
						}
						costPerMonth = cost * AverageHoursPerMonth
						break

					}
				}
			}
		}
	}

	return costPerMonth, nil
}
