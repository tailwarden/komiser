package rds

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func Clusters(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config rds.DescribeDBClustersInput
	resources := make([]Resource, 0)
	rdsClient := rds.NewFromConfig(*client.AWSClient)

	for {
		output, err := rdsClient.DescribeDBClusters(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, cluster := range output.DBClusters {
			tags := make([]Tag, 0)
			for _, tag := range cluster.TagList {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			var _clusterName string
			if cluster.DatabaseName == nil {
				_clusterName = *cluster.DBClusterIdentifier
			} else {
				_clusterName = *cluster.DatabaseName
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "RDS",
				Region:     client.AWSClient.Region,
				ResourceId: *cluster.DBClusterArn,
				Cost:       0,
				Name:       _clusterName,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/rds/home?region=%s#database:id=%s", client.AWSClient.Region, client.AWSClient.Region, *cluster.DBClusterIdentifier),
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
		"service":   "RDS",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
