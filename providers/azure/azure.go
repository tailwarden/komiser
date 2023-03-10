package digitalocean

import (
	"context"
	"log"

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
		networking.ApplicationGateways,
		networking.LoadBalancers,
		storage.Queues,
		storage.Databoxes,
		databases.Sql,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB, telemetry bool, analytics utils.Analytics) {
	for _, fetchResources := range listOfSupportedServices() {
		resources, err := fetchResources(ctx, client)
		if err != nil {
			log.Printf("[%s][Azure] %s", client.Name, err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).On("CONFLICT (resource_id) DO UPDATE").Set("cost = EXCLUDED.cost").Exec(context.Background())
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
