package elasticache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

func Clusters(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config elasticache.DescribeCacheClustersInput
	resources := make([]Resource, 0)
	elasticacheClient := elasticache.NewFromConfig(*client.AWSClient)

	pricingClient := pricing.NewFromConfig(*client.AWSClient)

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "Amazon ElastiCache")
	if err != nil {
		log.Warnln("Couldn't fetch Amazon ElastiCache cost and usage:", err)
	}

	for {
		output, err := elasticacheClient.DescribeCacheClusters(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, cluster := range output.CacheClusters {
			tagsResp, err := elasticacheClient.ListTagsForResource(ctx, &elasticache.ListTagsForResourceInput{
				ResourceName: cluster.ARN,
			})
			if err != nil {
				return resources, err
			}

			tags := make([]Tag, len(tagsResp.TagList))
			for i, t := range tagsResp.TagList {
				tags[i] = Tag{
					Key:   *t.Key,
					Value: *t.Value,
				}
			}

			monthlyCost := 0.0

			if *cluster.CacheClusterStatus != "available" {
				startOfMonth := utils.BeginningOfMonth(time.Now())
				hourlyUsage := 0
				if cluster.CacheClusterCreateTime.Before(startOfMonth) {
					hourlyUsage = int(time.Since(startOfMonth).Hours())
				} else {
					hourlyUsage = int(time.Since(*cluster.CacheClusterCreateTime).Hours())
				}

				pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
					ServiceCode: aws.String("AmazonElastiCache"),
					Filters: []types.Filter{
						{
							Field: aws.String("cacheEngine"),
							Value: aws.String(*cluster.Engine),
							Type:  types.FilterTypeTermMatch,
						},
						{
							Field: aws.String("instanceType"),
							Value: aws.String(string(*cluster.CacheNodeType)),
							Type:  types.FilterTypeTermMatch,
						},
						{
							Field: aws.String("location"),
							Value: aws.String("US East (N. Virginia)"),
							Type:  types.FilterTypeTermMatch,
						},
					},
					MaxResults: aws.Int32(1),
				})
				if err != nil {
					log.Warnf("Couldn't fetch pricing information for %s", *cluster.ARN)
				}

				hourlyCost := 0.0

				if pricingOutput != nil && len(pricingOutput.PriceList) > 0 {
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
					monthlyCost = float64(hourlyUsage) * hourlyCost
				}
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "ElastiCache",
				Region:     client.AWSClient.Region,
				ResourceId: *cluster.ARN,
				Name:       *cluster.CacheClusterId,
				Cost:       monthlyCost,
				CreatedAt:  *cluster.CacheClusterCreateTime,
				Tags:       tags,
				Metadata: map[string]string{
					"engine":    *cluster.Engine,
					"nodeType":  *cluster.CacheNodeType,
					"status":    *cluster.CacheClusterStatus,
					"clusterId": *cluster.CacheClusterId,
					"serviceCost": fmt.Sprint(serviceCost),
				},
				FetchedAt: time.Now(),
				Link:      fmt.Sprintf("https:/%s.console.aws.amazon.com/elasticache/home?region=%s#/%s/%s", client.AWSClient.Region, client.AWSClient.Region, *cluster.Engine, *cluster.CacheClusterId),
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
		"service":   "ElastiCache",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil

}
