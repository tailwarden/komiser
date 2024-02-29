package rds

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func ProxyEndpoints(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	config := rds.DescribeDBProxyEndpointsInput{
		MaxRecords: aws.Int32(100),
	}
	resources := make([]models.Resource, 0)
	rdsClient := rds.NewFromConfig(*client.AWSClient)
	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "RDS")
	if err != nil {
		log.Warnln("Couldn't fetch S3 cost and usage:", err)
	}
	for {
		output, err := rdsClient.DescribeDBProxyEndpoints(ctx, &config)
		if err != nil {
			return resources, fmt.Errorf("error describing DB proxy endpoints: %w", err)
		}

		for _, endpoint := range output.DBProxyEndpoints {

			if endpoint.DBProxyEndpointName == nil {
				log.Warn("DBProxyEndpointName is nil")
				continue
			}

			if endpoint.DBProxyEndpointArn == nil {
				log.Warn("DBProxyEndpointArn is nil")
				continue
			}

			_endpointName := *endpoint.DBProxyEndpointName

			// Fetch tags
			tagConfig := rds.ListTagsForResourceInput{
				ResourceName: endpoint.DBProxyEndpointArn,
			}

			tagOutput, err := rdsClient.ListTagsForResource(ctx, &tagConfig)
			if err != nil {
				return resources, fmt.Errorf("error listing tags for resource %s: %w", *endpoint.DBProxyEndpointArn, err)
			}

			tags := make([]models.Tag, 0)
			for _, tag := range tagOutput.TagList {
				if tag.Key != nil && tag.Value != nil {
					tags = append(tags, models.Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}
			}

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "RDS DB Proxy Endpoint",
				Region:     client.AWSClient.Region,
				ResourceId: *endpoint.DBProxyEndpointArn,
				Name:       _endpointName,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				Link: fmt.Sprintf("https://%s.console.aws.amazon.com/rds/home?region=%s#db-proxy-details:id=%s", client.AWSClient.Region, client.AWSClient.Region, *endpoint.DBProxyEndpointName),
			})
		}

		if output.Marker == nil || *output.Marker == "" {
			break
		}

		config.Marker = output.Marker
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "RDS DB Proxy Endpoint",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
