package linode

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/utils"

	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/linode/compute"
	"github.com/tailwarden/komiser/providers/linode/lkepool"
	"github.com/tailwarden/komiser/providers/linode/networking"
	"github.com/tailwarden/komiser/providers/linode/postgres"
	"github.com/tailwarden/komiser/providers/linode/sql"
	"github.com/tailwarden/komiser/providers/linode/storage"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		compute.LinodeInstancesAndInstanceDisks,
		compute.LKEClusters,
		storage.Volumes,
		storage.Databases,
		storage.Buckets,
		networking.NodeBalancers,
		networking.Firewalls,
		sql.Instances,
		postgres.Instances,
		lkepool.LKENodePools,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB, telemetry bool, analytics utils.Analytics, wp *providers.WorkerPool) {
	for _, fetchResources := range listOfSupportedServices() {
		fetchResources := fetchResources
		wp.SubmitTask(func() {
			resources, err := fetchResources(ctx, client)
			if err != nil {
				log.Printf("[%s][Linode] %s", client.Name, err)
			} else {
				for _, resource := range resources {
					_, err := db.NewInsert().Model(&resource).On("CONFLICT (resource_id) DO UPDATE").Set("cost = EXCLUDED.cost").Exec(context.Background())
					if err != nil {
						logrus.WithError(err).Errorf("db trigger failed")
					}
				}
				if telemetry {
					analytics.TrackEvent("discovered_resources", map[string]interface{}{
						"provider":  "Linode",
						"resources": len(resources),
					})
				}
			}
		})
	}
}
