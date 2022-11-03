package aws

import (
	"context"
	"log"

	. "github.com/mlabouardy/komiser/providers"
	. "github.com/mlabouardy/komiser/providers/digitalocean/droplets"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []FetchDataFunction {
	return []FetchDataFunction{
		Droplets,
	}
}

func FetchDigitalOceanData(ctx context.Context, client ProviderClient, db *bun.DB) {
	for _, function := range listOfSupportedServices() {
		resources, err := function(ctx, client)
		if err != nil {
			log.Printf("[%s][DigitalOcean] %s", client.Name, err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).Exec(context.Background())
			}
		}
	}
}
