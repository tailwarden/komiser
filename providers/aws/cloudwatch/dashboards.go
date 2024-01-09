package cloudwatch

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func Dashboards(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	cloudWatchClient := cloudwatch.NewFromConfig(*client.AWSClient)

	var nextToken *string
	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "AmazonCloudWatch")
	if err != nil {
		log.Warnln("Couldn't fetch AmazonCloudWatch cost and usage:", err)
	}

	for {
		input := &cloudwatch.ListDashboardsInput{
			NextToken: nextToken,
		}
		output, err := cloudWatchClient.ListDashboards(ctx, input)
		if err != nil {
			return resources, err
		}

		for index, dashboard := range output.DashboardEntries {
			outputTags, err := cloudWatchClient.ListTagsForResource(ctx, &cloudwatch.ListTagsForResourceInput{
				ResourceARN: dashboard.DashboardArn,
			})

			tags := make([]Tag, 0)

			if err == nil {
				for _, tag := range outputTags.Tags {
					tags = append(tags, Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}
			}

			cost := calculateDashboardCost(index + 1)

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "CloudWatch Dashboard",
				ResourceId: *dashboard.DashboardArn,
				Region:     client.AWSClient.Region,
				Name:       *dashboard.DashboardName,
				Cost:       cost,
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/cloudwatch/home?region=%s#dashboards:name=%s", client.AWSClient.Region, client.AWSClient.Region, *dashboard.DashboardName),
			})
		}

		if output.NextToken == nil {
			break
		}
		nextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "CloudWatch Dashboard",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}

func calculateDashboardCost(dashboardCount int) float64 {
	freeDashboards := 3
	cost := 0.0

	if dashboardCount > freeDashboards {
		cost = 3.0
	}
	return cost
}
