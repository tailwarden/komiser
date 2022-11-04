package aws

import (
	"context"
	"log"

	. "github.com/mlabouardy/komiser/providers"
	. "github.com/mlabouardy/komiser/providers/oci/compute"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []FetchDataFunction {
	return []FetchDataFunction{
		Instances,
	}
}

func FetchOciData(ctx context.Context, client ProviderClient, db *bun.DB) {
	for _, function := range listOfSupportedServices() {
		resources, err := function(ctx, client)
		if err != nil {
			log.Printf("[%s][OCI] %s", client.Name, err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).Exec(context.Background())
			}
		}
	}
}
