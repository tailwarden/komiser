package rds

import (
	"context"
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

		for _, instance := range output.DBProxies {
			// tags := make([]models.Tag, 0)
			// for _, tag := range instance.TagList {
			// 	tags = append(tags, models.Tag{
			// 		Key:   *tag.Key,
			// 		Value: *tag.Value,
			// 	})
			// }

			var _ProxyName string = *instance.DBProxyName
			startOfMonth := utils.BeginningOfMonth(time.Now())
			hourlyUsage := 0
			if (*instance.CreatedDate).Before(startOfMonth) {
				hourlyUsage = int(time.Since(startOfMonth).Hours())
			} else {
				hourlyUsage = int(time.Since(*instance.CreatedDate).Hours())
			}

			hourlyCost := 0.0
			monthlyCost := float64(hourlyUsage) * hourlyCost

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "RDS Instance",
				Region:     client.AWSClient.Region,
				ResourceId: *instance.DBProxyArn,
				Cost:       monthlyCost,
				Name:       _ProxyName,
				FetchedAt:  time.Now(),
				// Tags:       tags,
				// Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/rds/home?region=%s#database:id=%s", client.AWSClient.Region, client.AWSClient.Region, *instance.DBInstanceIdentifier),
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
