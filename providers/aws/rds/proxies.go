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
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

func Proxies(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config rds.DescribeDBProxiesInput
	resources := make([]models.Resource, 0)
	rdsClient := rds.NewFromConfig(*client.AWSClient)

	oldRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	client.AWSClient.Region = oldRegion

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "RDS")
	if err != nil {
		log.Warnln("Couldn't fetch S3 cost and usage:", err)
	}
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
				Service:    "RDS Proxy",
				Region:     client.AWSClient.Region,
				ResourceId: *proxy.DBProxyArn,
				Cost:       monthlyCost,
				Name:       _ProxyName,
				FetchedAt:  time.Now(),
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				Link: fmt.Sprintf("https:/%s.console.aws.amazon.com/rds/home?region=%s#proxies:id=%s", client.AWSClient.Region, client.AWSClient.Region, *proxy.DBProxyName),
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
		"service":   "RDS Proxy",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
