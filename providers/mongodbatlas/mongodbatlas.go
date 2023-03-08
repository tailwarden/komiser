package mongodbatlas

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/providers"
)

func FetchResources(ctx context.Context, client providers.ProviderClient) {
	// TODO fetch actual resources
	log.Info("Fetching MongoDBAtlas resources...")
}
