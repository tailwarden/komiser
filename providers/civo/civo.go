package civo

import (
	"context"
	"log"

	"github.com/civo/civogo"
	"github.com/mlabouardy/komiser/providers"
	"github.com/mlabouardy/komiser/providers/civo/compute"
	"github.com/mlabouardy/komiser/providers/civo/kubernetes"
	"github.com/mlabouardy/komiser/providers/civo/network"
	"github.com/mlabouardy/komiser/providers/civo/storage"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		compute.Instances,
		storage.Volumes,
		kubernetes.Clusters,
		network.Firewalls,
		network.Networks,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB) {
	for _, fetchResources := range listOfSupportedServices() {
		regions, err := client.CivoClient.ListRegions()
		if err != nil {
			log.Printf("[%s][Civo] %s", client.Name, err)
		}

		for _, region := range regions {
			clientWithRegion, err := civogo.NewClient(client.CivoClient.APIKey, region.Code)
			if err != nil {
				log.Printf("[%s][Civo] %s", client.Name, err)
			}

			client.CivoClient = clientWithRegion

			resources, err := fetchResources(ctx, client)
			if err != nil {
				log.Printf("[%s][Civo] %s", client.Name, err)
			} else {
				for _, resource := range resources {
					db.NewInsert().Model(&resource).Exec(context.Background())
				}
			}
		}
	}
}
