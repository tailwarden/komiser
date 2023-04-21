package systemsmanager

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func MaintenanceWindows(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	ssmClient := ssm.NewFromConfig(*client.AWSClient)

	input := &ssm.DescribeMaintenanceWindowsInput{}

	for {
		maintenanceWindows, err := ssmClient.DescribeMaintenanceWindows(ctx, input)
		if err != nil {
			return resources, err
		}

		for _, window := range maintenanceWindows.WindowIdentities {

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "SSM Maintenance Window",
				ResourceId: aws.ToString(window.WindowId),
				Name:       aws.ToString(window.Name),
				Region:     client.AWSClient.Region,
				Cost:       0,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/systems-manager/maintenance-windows/%s/details", client.AWSClient.Region, aws.ToString(window.WindowId)),
				Metadata: map[string]string{
					"Description": aws.ToString(window.Description),
				},
			})
		}

		if maintenanceWindows.NextToken == nil {
			break
		}

		input.NextToken = maintenanceWindows.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "SSM Maintenance Window",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
