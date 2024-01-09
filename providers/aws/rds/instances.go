package rds

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config rds.DescribeDBInstancesInput
	resources := make([]models.Resource, 0)
	rdsClient := rds.NewFromConfig(*client.AWSClient)

	oldRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	pricingClient := pricing.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = oldRegion
	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "Amazon Relational Database Service")
	if err != nil {
		log.Warnln("Couldn't fetch Amazon Relational Database Service cost and usage:", err)
	}

	for {
		output, err := rdsClient.DescribeDBInstances(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, instance := range output.DBInstances {
			tags := make([]models.Tag, 0)
			for _, tag := range instance.TagList {
				tags = append(tags, models.Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			var _instanceName string
			if instance.DBName == nil {
				_instanceName = *instance.DBInstanceIdentifier
			} else {
				_instanceName = *instance.DBName
			}

			startOfMonth := utils.BeginningOfMonth(time.Now())
			hourlyUsage := 0
			if (*instance.InstanceCreateTime).Before(startOfMonth) {
				hourlyUsage = int(time.Since(startOfMonth).Hours())
			} else {
				hourlyUsage = int(time.Since(*instance.InstanceCreateTime).Hours())
			}

			pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
				ServiceCode: aws.String("AmazonRDS"),
				Filters: []types.Filter{
					{
						Field: aws.String("instanceType"),
						Value: aws.String(*instance.DBInstanceClass),
						Type:  types.FilterTypeTermMatch,
					},
					{
						Field: aws.String("regionCode"),
						Value: aws.String(client.AWSClient.Region),
						Type:  types.FilterTypeTermMatch,
					},
					{
						Field: aws.String("databaseEngine"),
						Value: aws.String(*instance.Engine),
						Type:  types.FilterTypeTermMatch,
					},
				},
				MaxResults: aws.Int32(1),
			})
			if err != nil {
				log.Warnf("Couldn't fetch invocations metric for %s", _instanceName)
			}

			hourlyCost := 0.0
			if pricingOutput != nil && len(pricingOutput.PriceList) > 0 {
				log.Infof(`Raw pricingOutput.PriceList[0] : %s`, pricingOutput.PriceList[0])

				pricingResult := models.PricingResult{}
				err := json.Unmarshal([]byte(pricingOutput.PriceList[0]), &pricingResult)
				if err != nil {
					log.Fatalf("Failed to unmarshal JSON: %v", err)
				}

				for _, onDemand := range pricingResult.Terms.OnDemand {
					for _, priceDimension := range onDemand.PriceDimensions {
						hourlyCost, err = strconv.ParseFloat(priceDimension.PricePerUnit.USD, 64)
						if err != nil {
							log.Fatalf("Failed to parse hourly cost: %v", err)
						}
						break
					}
					break
				}

				//log.Printf("Hourly cost RDS: %f", hourlyCost)
			}

			monthlyCost := float64(hourlyUsage) * hourlyCost

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "RDS Instance",
				Region:     client.AWSClient.Region,
				ResourceId: *instance.DBInstanceArn,
				Cost:       monthlyCost,
				Metadata: map[string]string{
					"serviceCost":   fmt.Sprint(serviceCost),
					"engine":        *instance.Engine,
					"engineVersion": *instance.EngineVersion,
				},
				Name:      _instanceName,
				FetchedAt: time.Now(),
				Tags:      tags,
				Link:      fmt.Sprintf("https:/%s.console.aws.amazon.com/rds/home?region=%s#database:id=%s", client.AWSClient.Region, client.AWSClient.Region, *instance.DBInstanceIdentifier),
			})
		}

		if aws.ToString(output.Marker) == "" {
			break
		}

		config.Marker = output.Marker
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "RDS Instance",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
