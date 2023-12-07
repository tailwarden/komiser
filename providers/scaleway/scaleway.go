package scaleway

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/providers/scaleway/compute"
	"github.com/tailwarden/komiser/providers/scaleway/containers"
	"github.com/tailwarden/komiser/providers/scaleway/network"
	"github.com/tailwarden/komiser/providers/scaleway/storage"
	"github.com/tailwarden/komiser/utils"

	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/scaleway/serverless"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		serverless.Functions,
		serverless.ServerlessContainers,
		compute.Servers,
		containers.K8sClusters,
		containers.ContainerRegistries,
		storage.Databases,
		network.LoadBalancers,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB, telemetry bool, analytics utils.Analytics, wp *providers.WorkerPool) {
	wp.SubmitTask(func() {
		for _, fetchResources := range listOfSupportedServices() {
			fetchResources := fetchResources
			resources, err := fetchResources(ctx, client)
			if err != nil {
				log.Printf("[%s][Scaleway] %s", client.Name, err)
			} else {
				for _, resource := range resources {
					_, err := db.NewInsert().Model(&resource).On("CONFLICT (resource_id) DO UPDATE").Set("cost = EXCLUDED.cost").Exec(context.Background())
					if err != nil {
						logrus.WithError(err).Errorf("db trigger failed")
					}
				}
				if telemetry {
					analytics.TrackEvent("discovered_resources", map[string]interface{}{
						"provider":  "Scaleway",
						"resources": len(resources),
					})
				}
			}
		}
	})
}
