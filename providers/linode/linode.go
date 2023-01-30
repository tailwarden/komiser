package linode

import (
	"context"
	"log"

	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/linode/compute"
	"github.com/tailwarden/komiser/providers/linode/storage"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		compute.Linodes,
		storage.Volumes,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB) {
	for _, fetchResources := range listOfSupportedServices() {
		resources, err := fetchResources(ctx, client)
		if err != nil {
			log.Printf("[%s][Linode] %s", client.Name, err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).Exec(context.Background())
			}
		}
	}
}
