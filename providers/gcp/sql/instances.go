package sql

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/option"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	instancesClient, err := sqladmin.NewService(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create sql service")
		return resources, err
	}

	instances, err := instancesClient.Instances.List(client.GCPClient.Credentials.ProjectID).Do()
	if err != nil {
		if strings.Contains(err.Error(), "SERVICE_DISABLED") {
			logrus.Warn(err.Error())
			return resources, nil
		} else {
			logrus.WithError(err).Errorf("failed to list sql servers")
			return resources, err
		}
	}

	for _, sqlInstance := range instances.Items {
		resources = append(resources, models.Resource{
			Provider:  "GCP",
			Account:   client.Name,
			Service:   "SQL Instance",
			Name:      sqlInstance.Name,
			Region:    sqlInstance.Region,
			FetchedAt: time.Now(),
			Metadata: map[string]string{
				"databaseVersion": sqlInstance.DatabaseVersion,
			},
			Link: fmt.Sprintf("https://console.cloud.google.com/sql/instances/%s/overview?project=%s", sqlInstance.Name, client.GCPClient.Credentials.ProjectID),
		})
	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "SQL Instances",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
