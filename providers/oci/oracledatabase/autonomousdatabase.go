package oracledatabase

import (
	"context"
	"time"

	"github.com/oracle/oci-go-sdk/database"

	log "github.com/sirupsen/logrus"

	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func AutonomousDatabases(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	databaseClient, err := database.NewDatabaseClientWithConfigurationProvider(client.OciClient)
	if err != nil {
		return resources, err
	}

	tenancyOCID, err := client.OciClient.TenancyOCID()
	if err != nil {
		return resources, err
	}

	config := database.ListAutonomousDatabasesRequest{
		CompartmentId: &tenancyOCID,
	}

	output, err := databaseClient.ListAutonomousDatabases(context.Background(), config)
	if err != nil {
		return resources, err
	}

	for _, instance := range output.Items {
		tags := make([]Tag, 0)

		for key, value := range instance.FreeformTags {
			tags = append(tags, Tag{
				Key:   key,
				Value: value,
			})
		}

		region, err := client.OciClient.Region()
		if err != nil {
			return resources, err
		}

		resources = append(resources, Resource{
			Provider:   "OCI",
			Account:    client.Name,
			ResourceId: *instance.Id,
			Service:    "Autonomous Database",
			Region:     region,
			Name:       *instance.DisplayName,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "OCI",
		"account":   client.Name,
		"service":   "Autonomous Database",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
