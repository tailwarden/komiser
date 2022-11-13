package aws

import (
	"context"
	"log"

	"github.com/mlabouardy/komiser/providers"
	"github.com/mlabouardy/komiser/providers/digitalocean/droplets"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		droplets.Droplets,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB) {
	for _, fetchResources := range listOfSupportedServices() {
		resources, err := fetchResources(ctx, client)
		if err != nil {
			log.Printf("[%s][DigitalOcean] %s", client.Name, err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).Exec(context.Background())
			}
		}
	}
}
