package artifactregistry

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/option"

	ar "cloud.google.com/go/artifactregistry/apiv1"
	"cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
)

func ArtifactregistryPackages(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	arClient, err := ar.NewRESTClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create artifacts registry client")
		return resources, err
	}

	params := &artifactregistrypb.ListPackagesRequest{}
	pkgItr := arClient.ListPackages(ctx, params)

	for {
		pkg, err := pkgItr.Next()
		if err != nil {
			if strings.Contains(err.Error(), "SERVICE_DISABLED") {
				logrus.Warn(err.Error())
				return resources, nil
			} else {
				logrus.WithError(err).Errorf("failed to get next package")
				break
			}
		}
		if pkg == nil {
			break
		}

		resources = append(resources, models.Resource{
			Provider:   "GCP",
			Account:    client.Name,
			Service:    "Artifacts Registry Packages",
			ResourceId: pkg.Name,
			Name:       pkg.DisplayName,
			CreatedAt:  pkg.CreateTime.AsTime(),
			Cost:       0,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://console.cloud.google.com/artifacts/browse/%s?project=%s", client.GCPClient.Credentials.ProjectID, client.GCPClient.Credentials.ProjectID),
		})
	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "Artifacts Registry Packages",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
