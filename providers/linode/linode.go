package linode

import (
	"context"
	"log"

	"github.com/tailwarden/komiser/providers/linode/networking"
	"github.com/tailwarden/komiser/utils"

	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/linode/compute"
	"github.com/tailwarden/komiser/providers/linode/storage"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		compute.Linodes,
		compute.LKEClusters,
		storage.Volumes,
		storage.Databases,
		storage.Buckets,
		networking.NodeBalancers,
		networking.Firewalls,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB, telemetry bool, analytics utils.Analytics) {
	for _, fetchResources := range listOfSupportedServices() {
		resources, err := fetchResources(ctx, client)
		if err != nil {
			log.Printf("[%s][Linode] %s", client.Name, err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).On("CONFLICT (resource_id) DO UPDATE").Set("cost = EXCLUDED.cost").Exec(context.Background())
			}
			if telemetry {
				analytics.TrackEvent("discovered_resources", map[string]interface{}{
					"provider":  "Linode",
					"resources": len(resources),
				})
			}
		}
	}
}
