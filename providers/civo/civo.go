package civo

import (
	"context"
	"log"

	"github.com/civo/civogo"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/civo/compute"
	"github.com/tailwarden/komiser/providers/civo/kubernetes"
	"github.com/tailwarden/komiser/providers/civo/network"
	"github.com/tailwarden/komiser/providers/civo/storage"
	"github.com/tailwarden/komiser/utils"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		compute.Instances,
		storage.Volumes,
		storage.ObjectStores,
		storage.Databases,
		storage.DiskImages,
		kubernetes.Clusters,
		network.Firewalls,
		network.Networks,
		network.LoadBalancers,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB, telemetry bool, analytics utils.Analytics) {
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
					_, err := db.NewInsert().Model(&resource).On("CONFLICT (resource_id) DO UPDATE").Set("cost = EXCLUDED.cost").Exec(context.Background())
					if err != nil {
						logrus.WithError(err).Errorf("db trigger failed")
					}
				}
				if telemetry {
					analytics.TrackEvent("discovered_resources", map[string]interface{}{
						"provider":  "Azure",
						"resources": len(resources),
					})
				}
			}
		}
	}
}
