package gcp

import (
	"context"
	"github.com/tailwarden/komiser/providers/gcp/alloydb"

	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/gcp/appengine"
	"github.com/tailwarden/komiser/providers/gcp/artifactregistry"
	"github.com/tailwarden/komiser/providers/gcp/bigquery"
	certficate "github.com/tailwarden/komiser/providers/gcp/certificate"
	"github.com/tailwarden/komiser/providers/gcp/compute"
	"github.com/tailwarden/komiser/providers/gcp/container"
	"github.com/tailwarden/komiser/providers/gcp/firestore"
	"github.com/tailwarden/komiser/providers/gcp/function"
	"github.com/tailwarden/komiser/providers/gcp/gateway"
	"github.com/tailwarden/komiser/providers/gcp/iam"
	"github.com/tailwarden/komiser/providers/gcp/kms"
	"github.com/tailwarden/komiser/providers/gcp/redis"
	"github.com/tailwarden/komiser/providers/gcp/sql"
	"github.com/tailwarden/komiser/providers/gcp/storage"
	"github.com/tailwarden/komiser/utils"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		compute.Instances,
		compute.Disks,
		compute.Snapshots,
		storage.Buckets,
		bigquery.Tables,
		certficate.Certificates,
		iam.Roles,
		iam.ServiceAccounts,
		sql.Instances,
		redis.Instances,
		container.Clusters,
		kms.Keys,
		gateway.ApiGateways,
		function.Functions,
		alloydb.Instances,
		alloydb.Clusters,
		appengine.Services,
		artifactregistry.ArtifactregistryDockerImages,
		artifactregistry.ArtifactregistryPackages,
		artifactregistry.ArtifactregistryRepositories,
		firestore.Documents,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB, telemetry bool, analytics utils.Analytics, wp *providers.WorkerPool) {
	for _, fetchResources := range listOfSupportedServices() {
		fetchResources := fetchResources
		wp.SubmitTask(func() {
			resources, err := fetchResources(ctx, client)
			if err != nil {
				log.Printf("[%s][GCP] %s", client.Name, err)
			} else {
				for _, resource := range resources {
					_, err := db.NewInsert().Model(&resource).On("CONFLICT (resource_id) DO UPDATE").Set("cost = EXCLUDED.cost").Exec(context.Background())
					if err != nil {
						log.WithError(err).Errorf("db trigger failed")
					}

				}
				if telemetry {
					analytics.TrackEvent("discovered_resources", map[string]interface{}{
						"provider":  "GCP",
						"resources": len(resources),
					})
				}
			}
		})
	}
}
