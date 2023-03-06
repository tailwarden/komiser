package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Databases(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	rdbSvc := rdb.NewAPI(client.ScalewayClient)

	regions := []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}

	for _, region := range regions {
		outputInstances, err := rdbSvc.ListInstances(&rdb.ListInstancesRequest{
			Region: region,
		})
		if err != nil {
			return resources, err
		}

		for _, instance := range outputInstances.Instances {
			output, err := rdbSvc.ListDatabases(&rdb.ListDatabasesRequest{
				Region:     region,
				InstanceID: instance.ID,
			})
			if err != nil {
				return resources, err
			}

			for _, database := range output.Databases {
				resources = append(resources, models.Resource{
					Provider:   "Scaleway",
					Account:    client.Name,
					Service:    "Database",
					Region:     region.String(),
					ResourceId: database.Name,
					Cost:       0,
					Name:       database.Name,
					FetchedAt:  time.Now(),
					Link:       fmt.Sprintf("https://console.scaleway.com/rdb/instances/%s/%s", region.String(), database.Name),
				})
			}
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Scaleway",
		"account":   client.Name,
		"service":   "Database",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
