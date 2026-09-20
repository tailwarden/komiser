package ec2

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func NetworkInsightsAnalyses(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config ec2.DescribeNetworkInsightsAnalysesInput
	resources := make([]models.Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	for {
		output, err := ec2Client.DescribeNetworkInsightsAnalyses(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, analysis := range output.NetworkInsightsAnalyses {
			name := *analysis.NetworkInsightsAnalysisId
			tags := make([]models.Tag, 0)
			for _, tag := range analysis.Tags {
				if *tag.Key == "Name" {
					name = *tag.Value
				}
				tags = append(tags, models.Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			createdAt := time.Time{}
			if analysis.StartDate != nil {
				createdAt = *analysis.StartDate
			}

			networkPathFound := "false"
			if analysis.NetworkPathFound != nil && *analysis.NetworkPathFound {
				networkPathFound = "true"
			}

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Network Insights Analysis",
				Region:     client.AWSClient.Region,
				Name:       name,
				ResourceId: *analysis.NetworkInsightsAnalysisId,
				CreatedAt:  createdAt,
				FetchedAt:  time.Now(),
				Cost:       0,
				Tags:       tags,
				Link: fmt.Sprintf(
					"https:/%s.console.aws.amazon.com/vpc/home?region=%s#ReachabilityAnalyzerAnalysisDetails:networkInsightsAnalysisId=%s",
					client.AWSClient.Region,
					client.AWSClient.Region,
					*analysis.NetworkInsightsAnalysisId,
				),
				Metadata: map[string]string{
					"networkInsightsPathId": aws.ToString(analysis.NetworkInsightsPathId),
					"status":                string(analysis.Status),
					"networkPathFound":      networkPathFound,
				},
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
		"service":   "Network Insights Analysis",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
