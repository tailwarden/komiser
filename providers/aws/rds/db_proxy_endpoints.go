package rds

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func ProxyEndpoints(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	config := rds.DescribeDBProxyEndpointsInput{
		MaxRecords: aws.Int32(100),
	}
	resources := make([]models.Resource, 0)
	rdsClient := rds.NewFromConfig(*client.AWSClient)

	for {
		output, err := rdsClient.DescribeDBProxyEndpoints(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, endpoint := range output.DBProxyEndpoints {
			tags := make([]models.Tag, 0)
			//TODO (siddarthkay) : double check on tag retrieval
			_endpointName := *endpoint.DBProxyEndpointName

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "RDS DB Proxy Endpoint",
				Region:     client.AWSClient.Region,
				ResourceId: *endpoint.DBProxyEndpointArn,
				Name:       _endpointName,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/rds/home?region=%s#dbproxy:endpoint=%s", client.AWSClient.Region, client.AWSClient.Region, *endpoint.DBProxyEndpointName),
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
