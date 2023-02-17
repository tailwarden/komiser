package developerservices

import (
	"context"
	"time"

	"github.com/oracle/oci-go-sdk/functions"
	"github.com/tailwarden/komiser/providers"

	log "github.com/sirupsen/logrus"

	. "github.com/tailwarden/komiser/models"
)

func Functions(ctx context.Context, applicationId *string, client providers.ProviderClient, functionsManagementClient functions.FunctionsManagementClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	listFunctionsRequest := functions.ListFunctionsRequest{
		ApplicationId: applicationId,
	}

	output, err := functionsManagementClient.ListFunctions(context.Background(), listFunctionsRequest)
	if err != nil {
		return resources, err
	}

	for _, function := range output.Items {
		tags := make([]Tag, 0)

		for key, value := range function.FreeformTags {
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
			ResourceId: *function.Id,
			Service:    "Function",
			Region:     region,
			Name:       *function.DisplayName,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "OCI",
		"account":   client.Name,
		"service":   "Function",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
