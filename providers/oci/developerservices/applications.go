package developerservices

import (
	"context"
	"time"

	"github.com/oracle/oci-go-sdk/functions"

	log "github.com/sirupsen/logrus"

	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Applications(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	functionsManagementClient, err := functions.NewFunctionsManagementClientWithConfigurationProvider(client.OciClient)
	if err != nil {
		return resources, err
	}

	tenancyOCID, err := client.OciClient.TenancyOCID()
	if err != nil {
		return resources, err
	}

	listApplicationsRequest := functions.ListApplicationsRequest{
		CompartmentId: &tenancyOCID,
	}

	output, err := functionsManagementClient.ListApplications(context.Background(), listApplicationsRequest)
	if err != nil {
		return resources, err
	}

	for _, application := range output.Items {
		tags := make([]Tag, 0)

		for key, value := range application.FreeformTags {
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

		// add functions to resources [containing applications] too, as functions are sub-resources to applications.
		allFunctions, err := Functions(ctx, application.Id, client, functionsManagementClient)
		if err != nil {
			return resources, err
		}
		if len(allFunctions) > 0 {
			resources = append(resources, allFunctions...)
		}

		resources = append(resources, Resource{
			Provider:   "OCI",
			Account:    client.Name,
			ResourceId: *application.Id,
			Service:    "Application",
			Region:     region,
			Name:       *application.DisplayName,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "OCI",
		"account":   client.Name,
		"service":   "Application",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
