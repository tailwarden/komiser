package scaleway

import (
	"context"
	"github.com/tailwarden/komiser/providers/scaleway/compute"
	"github.com/tailwarden/komiser/providers/scaleway/containers"
	"github.com/tailwarden/komiser/providers/scaleway/manageddbs"
	"github.com/tailwarden/komiser/providers/scaleway/network"
	"log"

	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/scaleway/serverless"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		serverless.Functions,
		compute.Instances,
		containers.K8sClusters,
		containers.ContainerRegistries,
		manageddbs.PostgresAndMySQLs,
		network.LoadBalancers,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB) {
	for _, fetchResources := range listOfSupportedServices() {
		resources, err := fetchResources(ctx, client)
		if err != nil {
			log.Printf("[%s][Scaleway] %s", client.Name, err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).Exec(context.Background())
			}
		}
	}
}
