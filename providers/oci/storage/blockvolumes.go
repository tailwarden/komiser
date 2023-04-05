package storage

import (
	"context"
	"time"

	"github.com/oracle/oci-go-sdk/core"

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

		region, err := client.OciClient.Region()
		if err != nil {
			return resources, err
		}

	//Calculate cost 
	volumeRequest := core.GetVolumeRequest{
		VolumeId: volume.Id,
	}
	volumeResponse, err := blockStorageClient.GetVolume(context.Background(), volumeRequest)
	if err != nil {
		return resources, err
	}
	perGBMonth := 0.0255
	
	cost := float64(*volumeResponse.Volume.SizeInGBs) * perGBMonth


		resources = append(resources, Resource{
			Provider:   "OCI",
			Account:    client.Name,
			ResourceId: *volume.Id,
			Service:    "Block Volume",
			Region:     region,
			Name:       *volume.DisplayName,
			Cost:       cost,
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
