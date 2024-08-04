package function

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/cloudfunctions/v2"
	"google.golang.org/api/option"
)

func Functions(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	cloudFunctionsService, err := cloudfunctions.NewService(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		log.WithError(err).Errorf("failed to create Cloud Functions service")
		return resources, err
	}

	functions, err := cloudFunctionsService.Projects.Locations.Functions.List(
		"projects/" + client.GCPClient.Credentials.ProjectID + "/locations/-",
	).Do()
	if err != nil {
		if strings.Contains(err.Error(), "SERVICE_DISABLED") {
			log.Warn(err.Error())
			return resources, nil
		} else {
			log.WithError(err).Errorf("failed to list Cloud Functions")
			return resources, err
		}
	}

	for _, function := range functions.Functions {
		fmt.Printf("%+v\n", function)

		re := regexp.MustCompile(`projects\/.*?\/locations\/(.+?)\/functions\/(.+)`)
		match := re.FindStringSubmatch(function.Name)

		generation := "gen1"
		functionRegion := ""
		functionName := function.Name

		if function.Environment == "GEN_2" {
			generation = "gen2"
		}
		if len(match) == 3 {
			functionRegion = match[1]
			functionName = match[2]

		}

		resources = append(resources, models.Resource{
			Provider:   "GCP",
			Account:    client.Name,
			Service:    "Cloud Functions",
			ResourceId: function.Name,
			Region:     functionRegion,
			Name:       functionName,
			Metadata: map[string]string{
				"Version":       function.Environment,
				"Last Modified": function.UpdateTime,
			},
			FetchedAt: time.Now(),
			Link:      fmt.Sprintf("https://console.cloud.google.com/functions/details/%s/%s?env=%sproject=%s", functionRegion, functionName, generation, client.GCPClient.Credentials.ProjectID),
		})

	}

	log.WithFields(log.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "Cloud Functions",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
