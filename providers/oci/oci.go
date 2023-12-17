package oci

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/providers/oci/developerservices"
	"github.com/tailwarden/komiser/providers/oci/iam"
	"github.com/tailwarden/komiser/providers/oci/oracledatabase"
	"github.com/tailwarden/komiser/providers/oci/storage"
	"github.com/tailwarden/komiser/utils"

	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/oci/compute"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		compute.Instances,
		iam.Policies,
		oracledatabase.AutonomousDatabases,
		storage.Buckets,
		storage.BlockVolumes,
		developerservices.Applications,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB, telemetry bool, analytics utils.Analytics, wp *providers.WorkerPool) {
	for _, fetchResources := range listOfSupportedServices() {
		fetchResources := fetchResources
		wp.SubmitTask(func() {
			resources, err := fetchResources(ctx, client)
			if err != nil {
				log.Printf("[%s][OCI] %s", client.Name, err)
			} else {
				for _, resource := range resources {
					_, err = db.NewInsert().Model(&resource).On("CONFLICT (resource_id) DO UPDATE").Set("cost = EXCLUDED.cost").Exec(context.Background())
					if err != nil {
						logrus.WithError(err).Errorf("db trigger failed")
					}
				}
				if telemetry {
					analytics.TrackEvent("discovered_resources", map[string]interface{}{
						"provider":  "OCI",
						"resources": len(resources),
					})
				}
			}
		})
	}
}
