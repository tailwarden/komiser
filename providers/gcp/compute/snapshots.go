package compute

import (
	"context"
	"fmt"
	"strings"
	"time"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils/gcpcomputepricing"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Snapshots(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	snapshotsClient, err := compute.NewSnapshotsRESTClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create snapshots client")
		return resources, err
	}

	req := &computepb.ListSnapshotsRequest{
		Project: client.GCPClient.Credentials.ProjectID,
	}

	snapshots := snapshotsClient.List(ctx, req)

	actualPricing, err := gcpcomputepricing.Fetch()
	if err != nil {
		logrus.WithError(err).Errorf("failed to fetch actual GCP snapshots pricing")
		return resources, err
	}

	for {
		snapshot, err := snapshots.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			if strings.Contains(err.Error(), "SERVICE_DISABLED") {
				logrus.Warn(err.Error())
				return resources, nil
			} else {
				logrus.WithError(err).Errorf("failed to list snapshots")
				return resources, err
			}
		}

		tags := make([]models.Tag, 0)
		if snapshot.Labels != nil {
			for key, value := range snapshot.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}
		}

		var region string
		locations := snapshot.GetStorageLocations()
		if len(locations) > 0 {
			region = locations[0]
		}

		cost, err := gcpcomputepricing.CalculateSnapshotCost(ctx, client, gcpcomputepricing.CalculateSnapshotCostData{
			StorageBytes:      snapshot.GetStorageBytes(),
			CreationTimestamp: snapshot.GetCreationTimestamp(),
			Region:            region,
			Project:           client.GCPClient.Credentials.ProjectID,
			// Zone:     zone,
			Pricing: actualPricing,
		})
		if err != nil {
			logrus.WithError(err).Errorf("failed to calculate disk cost")
			return resources, err
		}

		resources = append(resources, models.Resource{
			Provider:   "GCP",
			Account:    client.Name,
			Service:    "Compute Disk Snapshot",
			ResourceId: fmt.Sprintf("%d", snapshot.GetId()),
			Region:     region,
			Name:       snapshot.GetName(),
			Cost:       cost,
			FetchedAt:  time.Now(),
			Tags:       tags,
			Link:       fmt.Sprintf("https://console.cloud.google.com/compute/snapshotsDetail/projects/%s/global/snapshots/%s?project=%s", client.GCPClient.Credentials.ProjectID, snapshot.GetName(), client.GCPClient.Credentials.ProjectID),
		})
	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "Compute Disk Snapshot",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
