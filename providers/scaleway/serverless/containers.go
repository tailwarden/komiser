package serverless

import (
	"context"
	"fmt"
	"time"

	container "github.com/scaleway/scaleway-sdk-go/api/container/v1beta1"

	"github.com/scaleway/scaleway-sdk-go/scw"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func ServerlessContainers(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	containerSvc := container.NewAPI(client.ScalewayClient)

	regions := []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}

	for _, region := range regions {
		output, err := containerSvc.ListNamespaces(&container.ListNamespacesRequest{
			Region: region,
		})
		if err != nil {
			return resources, err
		}
		for _, namespace := range output.Namespaces {
			resources = append(resources, models.Resource{
				Provider:   "Scaleway",
				Account:    client.Name,
				Service:    "ServerlessContainer",
				Region:     namespace.Region.String(),
				ResourceId: namespace.ID,
				Cost:       0,
				Name:       namespace.Name,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://console.scaleway.com/containers/namespaces/%s/%s/containers", namespace.Region.String(), namespace.ID),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Scaleway",
		"account":   client.Name,
		"service":   "ServerlessContainer",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
