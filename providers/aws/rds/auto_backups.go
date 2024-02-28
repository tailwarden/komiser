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

func AutoBackups(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config rds.DescribeDBInstanceAutomatedBackupsInput
	resources := make([]models.Resource, 0)
	rdsClient := rds.NewFromConfig(*client.AWSClient)
	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "RDS")
	if err != nil {
		log.Warnln("Couldn't fetch S3 cost and usage:", err)
	}
	for {
		output, err := rdsClient.DescribeDBInstanceAutomatedBackups(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, backup := range output.DBInstanceAutomatedBackups {

			_backupName := *backup.DBInstanceIdentifier

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "RDS Backup",
				Region:     client.AWSClient.Region,
				ResourceId: *backup.DBInstanceArn,
				Name:       _backupName,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/rds/home?region=%s#dbinstance:id=%s", client.AWSClient.Region, client.AWSClient.Region, *backup.DBInstanceIdentifier),
				Metadata: map[string]string{
					"serviceCost":   fmt.Sprint(serviceCost),
					"Engine":        *backup.Engine,
					"EngineVersion": *backup.EngineVersion,
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
		"service":   "RDS Backup",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
