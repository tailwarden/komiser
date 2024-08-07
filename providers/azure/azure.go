package azure

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/azure/compute"
	"github.com/tailwarden/komiser/providers/azure/databases"
	"github.com/tailwarden/komiser/providers/azure/networking"
	"github.com/tailwarden/komiser/providers/azure/storage"
	"github.com/tailwarden/komiser/utils"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		compute.Disks,
		compute.Images,
		compute.VirtualMachines,
		compute.Snapshots,
		networking.ApplicationGateways,
		networking.LoadBalancers,
		networking.Firewalls,
		networking.LocalNetworkGateways,
		storage.Queues,
		storage.Tables,
		storage.Databoxes,
		databases.Sql,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB, telemetry bool, analytics utils.Analytics, wp *providers.WorkerPool) {
	for _, fetchResources := range listOfSupportedServices() {
		fetchResources := fetchResources
		wp.SubmitTask(func() {
			resources, err := fetchResources(ctx, client)
			if err != nil {
				log.Printf("[%s][Azure] %s", client.Name, err)
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
		})
	}
}
