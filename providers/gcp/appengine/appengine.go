package appengine

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/option"

	"cloud.google.com/go/appengine/apiv1"
	"cloud.google.com/go/appengine/apiv1/appenginepb"
)

func Appengine(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	appClient, err := appengine.NewApplicationsClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create disks client")
		return resources, err
	}

	params := appenginepb.GetApplicationRequest{}
	apps, err := appClient.GetApplication(ctx, &params)
	if err != nil {
		logrus.WithError(err).Errorf("failed to get app engine")
		return resources, err
	}
	resources = append(resources, models.Resource{
		Provider:   "GCP",
		Account:    client.Name,
		Service:    "App Engine",
		ResourceId: apps.Id,
		Name:       apps.Name,
		Cost:       0,
		FetchedAt:  time.Now(),
		Link:       fmt.Sprintf("https://console.cloud.google.com/security/ccm/certificates/details/global/name/%s?project=%s", client.GCPClient.Credentials.ProjectID),
	})


	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "App Engine",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
