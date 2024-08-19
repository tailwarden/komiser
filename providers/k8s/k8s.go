package k8s

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/k8s/core"
	"github.com/tailwarden/komiser/utils"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		core.Pods,
		core.Services,
		core.Deployments,
		core.Ingress,
		core.PersistentVolumes,
		core.PersistentVolumeClaims,
		core.ServiceAccounts,
		core.Nodes,
		core.Namespaces,
		core.DaemonSets,
		core.StatefulSets,
		core.Jobs,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB, telemetry bool, analytics utils.Analytics, wp *providers.WorkerPool) {
	for _, fetchResources := range listOfSupportedServices() {
		fetchResources := fetchResources
		wp.SubmitTask(func() {
			resources, err := fetchResources(ctx, client)
			if err != nil {
				log.Printf("[%s][K8s] %s", client.Name, err)
			} else {
				for _, resource := range resources {
					_, err = db.NewInsert().Model(&resource).On("CONFLICT (resource_id) DO UPDATE").Set("cost = EXCLUDED.cost, relations=EXCLUDED.relations").Exec(context.Background())
					if err != nil {
						logrus.WithError(err).Errorf("db trigger failed")
					}
				}
				if telemetry {
					analytics.TrackEvent("discovered_resources", map[string]interface{}{
						"provider":  "K8s",
						"resources": len(resources),
					})
				}
			}
		})
	}
}
