package compute

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/oracle/oci-go-sdk/core"
	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Instances(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	computeClient, err := core.NewComputeClientWithConfigurationProvider(client.OciClient)
	if err != nil {
		return resources, err
	}

	tenancyOCID, err := client.OciClient.TenancyOCID()
	if err != nil {
		return resources, err
	}

	config := core.ListInstancesRequest{
		CompartmentId: &tenancyOCID,
	}

	output, err := computeClient.ListInstances(context.Background(), config)
	if err != nil {
		return resources, err
	}

	for _, instance := range output.Items {
		tags := make([]Tag, 0)

		for key, value := range instance.FreeformTags {
			tags = append(tags, Tag{
				Key:   key,
				Value: value,
			})
		}

		resources = append(resources, Resource{
			Provider:   "OCI",
			Account:    client.Name,
			ResourceId: *instance.Id,
			Service:    "VM",
			Region:     *instance.Region,
			Name:       *instance.DisplayName,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "OCI",
		"account":   client.Name,
		"service":   "Compute VM",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
