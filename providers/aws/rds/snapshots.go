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

func Snapshots(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config rds.DescribeDBSnapshotsInput
	resources := make([]models.Resource, 0)
	rdsClient := rds.NewFromConfig(*client.AWSClient)

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "RDS")
	if err != nil {
		log.Warnln("Couldn't fetch S3 cost and usage:", err)
	}
	for {
		output, err := rdsClient.DescribeDBSnapshots(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, snapshot := range output.DBSnapshots {
			tags := make([]models.Tag, 0)
			for _, tag := range snapshot.TagList {
				tags = append(tags, models.Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			_snapshotName := *snapshot.DBSnapshotIdentifier

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "RDS Snapshot",
				Region:     client.AWSClient.Region,
				ResourceId: *snapshot.DBSnapshotArn,
				Name:       _snapshotName,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				Link: fmt.Sprintf("https://%s.console.aws.amazon.com/rds/home?region=%s#snapshot:id=%s", client.AWSClient.Region, client.AWSClient.Region, *snapshot.DBSnapshotIdentifier),
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
		"service":   "RDS Snapshot",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
