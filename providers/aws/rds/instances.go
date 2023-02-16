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

func Instances(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config rds.DescribeDBInstancesInput
	resources := make([]Resource, 0)
	rdsClient := rds.NewFromConfig(*client.AWSClient)

	for {
		output, err := rdsClient.DescribeDBInstances(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, instance := range output.DBInstances {
			tags := make([]Tag, 0)
			for _, tag := range instance.TagList {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			var _instanceName string
			if instance.DBName == nil {
				_instanceName = *instance.DBInstanceIdentifier
			} else {
				_instanceName = *instance.DBName
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "RDS Instances",
				Region:     client.AWSClient.Region,
				ResourceId: *instance.DBInstanceArn,
				Cost:       0,
				Name:       _instanceName,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/rds/home?region=%s#database:id=%s", client.AWSClient.Region, client.AWSClient.Region, *instance.DBInstanceIdentifier),
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
		"service":   "RDS Instances",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
