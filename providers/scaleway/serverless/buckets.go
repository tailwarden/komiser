package serverless

import (
	"context"
	"fmt"
	"time"

	function "github.com/scaleway/scaleway-sdk-go/api/function/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Functions(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	functionSvc := function.NewAPI(client.ScalewayClient)

	regions := []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}

	for _, region := range regions {
		output, err := functionSvc.ListFunctions(&function.ListFunctionsRequest{
			Region: region,
		})
		if err != nil {
			return resources, err
		}

		for _, function := range output.Functions {
			resources = append(resources, models.Resource{
				Provider:   "Scaleway",
				Account:    client.Name,
				Service:    "Function",
				Region:     function.Region.String(),
				ResourceId: function.ID,
				Cost:       0,
				Name:       function.Name,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://console.scaleway.com/functions/namespaces/%s/%s/functions", function.Region.String(), function.ID),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Scaleway",
		"account":   client.Name,
		"service":   "Function",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
