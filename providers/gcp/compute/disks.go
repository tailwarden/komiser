package compute

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils"
	"github.com/tailwarden/komiser/utils/gcpcomputepricing"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
)

func Disks(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	disksClient, err := compute.NewDisksRESTClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create disks client")
		return resources, err
	}

	req := &computepb.AggregatedListDisksRequest{
		Project: client.GCPClient.Credentials.ProjectID,
	}
	disks := disksClient.AggregatedList(ctx, req)

	actualPricing, err := gcpcomputepricing.Fetch()
	if err != nil {
		logrus.WithError(err).Errorf("failed to fetch actual GCP disks pricing")
		return resources, err
	}

	for {
		disksListPair, err := disks.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			if strings.Contains(err.Error(), "SERVICE_DISABLED") {
				logrus.Warn(err.Error())
				return resources, nil
			} else {
				logrus.WithError(err).Errorf("failed to list disks")
				return resources, err
			}
		}
		if len(disksListPair.Value.Disks) == 0 {
			continue
		}

		for _, disk := range disksListPair.Value.Disks {
			tags := make([]models.Tag, 0)
			if disk.Labels != nil {
				for key, value := range disk.Labels {
					tags = append(tags, models.Tag{
						Key:   key,
						Value: value,
					})
				}
			}

			zone := utils.GcpExtractZoneFromURL(disk.GetZone())
			size := disk.GetSizeGb()

			cost, err := gcpcomputepricing.CalculateDiskCost(ctx, client, gcpcomputepricing.CalculateDiskCostData{
				DiskType:          disk.GetType(),
				Size:              size,
				CreationTimestamp: disk.GetCreationTimestamp(),
				Project:           client.GCPClient.Credentials.ProjectID,
				Zone:              zone,
				Pricing:           actualPricing,
			})
			if err != nil {
				logrus.WithError(err).Errorf("failed to calculate disk cost")
			}

			resources = append(resources, models.Resource{
				Provider:   "GCP",
				Account:    client.Name,
				Service:    "Compute Disk",
				ResourceId: fmt.Sprintf("%d", disk.GetId()),
				Region:     zone,
				Name:       disk.GetName(),
				Cost:       cost,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https://console.cloud.google.com/compute/disksDetail/zones/%s/disks/%s?project=%s", zone, disk.GetName(), client.GCPClient.Credentials.ProjectID),
			})
		}
	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "Compute Disk",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
