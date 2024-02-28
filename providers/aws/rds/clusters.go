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
)

func Clusters(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config rds.DescribeDBClustersInput
	resources := make([]models.Resource, 0)
	rdsClient := rds.NewFromConfig(*client.AWSClient)
	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "RDS")
	if err != nil {
		log.Warnln("Couldn't fetch S3 cost and usage:", err)
	}
	for {
		output, err := rdsClient.DescribeDBClusters(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, cluster := range output.DBClusters {
			tags := make([]models.Tag, 0)
			for _, tag := range cluster.TagList {
				tags = append(tags, models.Tag{
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

			resources = append(resources, models.Resource{
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
				Metadata: map[string]string{
					"Engine":        *cluster.Engine,
					"EngineVersion": *cluster.EngineVersion,
					"serviceCost":   fmt.Sprint(serviceCost),
				},
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
