package containers

import (
	"context"
	"fmt"
	"github.com/scaleway/scaleway-sdk-go/api/registry/v1"
	"time"

	"github.com/scaleway/scaleway-sdk-go/scw"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func ContainerRegistries(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	crSvc := registry.NewAPI(client.ScalewayClient)

	regions := []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}

	for _, region := range regions {
		output, err := crSvc.ListNamespaces(&registry.ListNamespacesRequest{
			Region: region,
		})
		if err != nil {
			return resources, err
		}
		for _, namespace := range output.Namespaces {
			resources = append(resources, models.Resource{
				Provider:   "Scaleway",
				Account:    client.Name,
				Service:    "ContainerRegistry",
				Region:     namespace.Region.String(),
				ResourceId: namespace.ID,
				Cost:       0,
				Name:       namespace.Name,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://console.scaleway.com/registry/namespaces/%s/%s", namespace.Region.String(), namespace.ID),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Scaleway",
		"account":   client.Name,
		"service":   "ContainerRegistry",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
