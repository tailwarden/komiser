package rds

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils"
)

func Proxies(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config rds.DescribeDBProxiesInput
	resources := make([]models.Resource, 0)
	rdsClient := rds.NewFromConfig(*client.AWSClient)

	oldRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	client.AWSClient.Region = oldRegion

	for {
		output, err := rdsClient.DescribeDBProxies(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, proxy := range output.DBProxies {
			var _ProxyName string = *proxy.DBProxyName
			startOfMonth := utils.BeginningOfMonth(time.Now())
			hourlyUsage := 0
			if (*proxy.CreatedDate).Before(startOfMonth) {
				hourlyUsage = int(time.Since(startOfMonth).Hours())
			} else {
				hourlyUsage = int(time.Since(*proxy.CreatedDate).Hours())
			}

			hourlyCost := 0.0
			monthlyCost := float64(hourlyUsage) * hourlyCost

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "RDS Instance",
				Region:     client.AWSClient.Region,
				ResourceId: *proxy.DBProxyArn,
				Cost:       monthlyCost,
				Name:       _ProxyName,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/rds/home?region=%s#proxies:id=%s", client.AWSClient.Region, client.AWSClient.Region, *proxy.DBProxyName),
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
