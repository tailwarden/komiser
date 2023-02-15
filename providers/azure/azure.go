package digitalocean

import (
	"context"
	"log"

	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/azure/compute"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		compute.Disks,
		compute.Images,
		compute.VirtualMachines,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB) {
	for _, fetchResources := range listOfSupportedServices() {
		resources, err := fetchResources(ctx, client)
		if err != nil {
			log.Printf("[%s][Azure] %s", client.Name, err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).Exec(context.Background())
			}
		}
	}
}
