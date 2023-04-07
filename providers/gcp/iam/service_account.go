package iam

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/iterator"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"cloud.google.com/go/iam"
)

func ServiceAccounts(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	serviceClient, err := iamcredentials.GenerateAccessTokenRequest
	if err != nil {
		logrus.WithError(err).Errorf("failed to create compute client")
		return resources, err
	}

	req := &computepb.AggregatedListInstancesRequest{
		Project: client.GCPClient.Credentials.ProjectID,
	}
	instances := instancesClient.AggregatedList(ctx, req)

	for {
		instanceListPair, err := instances.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logrus.WithError(err).Errorf("failed to list instances")
			return resources, err
		}
		if len(instanceListPair.Value.Instances) == 0 {
			continue
		}

		for _, instance := range instanceListPair.Value.Instances {
			tags := make([]models.Tag, 0)
			if instance.Labels != nil {
				for key, value := range instance.Labels {
					tags = append(tags, models.Tag{
						Key:   key,
						Value: value,
					})
				}
			}

			zone := extractZoneFromURL(instance.GetZone())

			resources = append(resources, models.Resource{
				Provider:   "GCP",
				Account:    client.Name,
				Service:    "VM Instance",
				ResourceId: fmt.Sprintf("%d", instance.GetId()),
				Region:     zone,
				Name:       instance.GetName(),
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https://console.cloud.google.com/compute/instancesDetail/zones/%s/instances/%s?project=%s", zone, instance.GetName(), client.GCPClient.Credentials.ProjectID),
			})
		}
	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "Compute Engine",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}

func extractZoneFromURL(url string) string {
	return url[strings.LastIndex(url, "/")+1:]
}
