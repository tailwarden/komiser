package storage

import (
	"context"
	"github.com/oracle/oci-go-sdk/core"
	"time"

	log "github.com/sirupsen/logrus"

	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func BlockVolumes(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	blockStorageClient, err := core.NewBlockstorageClientWithConfigurationProvider(client.OciClient)
	if err != nil {
		return resources, err
	}

	tenancyOCID, err := client.OciClient.TenancyOCID()
	if err != nil {
		return resources, err
	}

	config := core.ListVolumesRequest{
		CompartmentId: &tenancyOCID,
	}

	output, err := blockStorageClient.ListVolumes(context.Background(), config)
	if err != nil {
		return resources, err
	}

	for _, volume := range output.Items {
		tags := make([]Tag, 0)

		for key, value := range volume.FreeformTags {
			tags = append(tags, Tag{
				Key:   key,
				Value: value,
			})
		}

		region, err1 := client.OciClient.Region()
		if err1 != nil {
			return resources, err1
		}

		resources = append(resources, Resource{
			Provider:   "OCI",
			Account:    client.Name,
			ResourceId: *volume.Id,
			Service:    "Block Volume",
			Region:     region,
			Name:       *volume.DisplayName,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "OCI",
		"account":   client.Name,
		"service":   "Block Volume",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
