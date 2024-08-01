package appengine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/option"

	appengine "cloud.google.com/go/appengine/apiv1"
	"cloud.google.com/go/appengine/apiv1/appenginepb"
)

func Services(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	appClient, err := appengine.NewServicesClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create app engine client")
		return resources, err
	}

	params := appenginepb.ListServicesRequest{}
	svcs := appClient.ListServices(ctx, &params)

	for {
		svc, err := svcs.Next()
		if err != nil {
			if strings.Contains(err.Error(), "SERVICE_DISABLED") {
				logrus.Warn(err.Error())
				return resources, nil
			} else {
				logrus.WithError(err).Errorf("failed to get app engine")
				break
			}
		}
		if svc == nil {
			break
		}

		resources = append(resources, models.Resource{
			Provider:   "GCP",
			Account:    client.Name,
			Service:    "App Engine",
			ResourceId: svc.Id,
			Name:       svc.Name,
			Cost:       0,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://console.cloud.google.com/appengine/services?project=%s", client.GCPClient.Credentials.ProjectID),
		})
	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "App Engine",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
