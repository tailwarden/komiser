package iam

import (
	"context"
	"time"

	"github.com/oracle/oci-go-sdk/identity"

	log "github.com/sirupsen/logrus"

	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Policies(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	identityClient, err := identity.NewIdentityClientWithConfigurationProvider(client.OciClient)
	if err != nil {
		return resources, err
	}

	tenancyOCID, err := client.OciClient.TenancyOCID()
	if err != nil {
		return resources, err
	}

	config := identity.ListPoliciesRequest{
		CompartmentId: &tenancyOCID,
	}

	output, err := identityClient.ListPolicies(context.Background(), config)
	if err != nil {
		return resources, err
	}

	for _, policy := range output.Items {
		tags := make([]Tag, 0)

		for key, value := range policy.FreeformTags {
			tags = append(tags, Tag{
				Key:   key,
				Value: value,
			})
		}

		// extract region from client
		region, err := client.OciClient.Region()
		if err != nil {
			return resources, err
		}

		resources = append(resources, Resource{
			Provider:   "OCI",
			Account:    client.Name,
			ResourceId: *policy.Id,
			Service:    "Identity Policy",
			Region:     region,
			Name:       *policy.Name,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "OCI",
		"account":   client.Name,
		"service":   "Identity Policy",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
