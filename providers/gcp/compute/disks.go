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
			logrus.WithError(err).Errorf("failed to list disks")
			return resources, err
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

			cost, err := calculateDiskCost(ctx, client, calculateDiskCostData{
				diskType: disk.GetType(),
				size:     size,
				project:  client.GCPClient.Credentials.ProjectID,
				zone:     zone,
				pricing:  actualPricing,
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
		"service":   "Compute Engine",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}

type calculateDiskCostData struct {
	diskType string
	size     int64
	project  string
	zone     string
	pricing  *gcpcomputepricing.Pricing
}

func calculateDiskCost(ctx context.Context, client providers.ProviderClient, data calculateDiskCostData) (float64, error) {
	diskTypeClient, err := compute.NewDiskTypesRESTClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		return 0, err
	}

	dtS := strings.Split(data.diskType, "/")
	dt, err := diskTypeClient.Get(ctx, &computepb.GetDiskTypeRequest{
		DiskType: dtS[len(dtS)-1],
		Project:  data.project,
		Zone:     data.zone,
	})
	if err != nil {
		return 0, err
	}

	var opts = gcpcomputepricing.Opts{
		Region:   utils.GcpGetRegionFromZone(data.zone),
		DiskType: dtS[len(dtS)-1],
		DiskSize: uint64(data.size),
	}
	var cost float64
	if dt.Name != nil {
		switch {
		case strings.Contains(strings.ToLower(*dt.Name), strings.ToLower(gcpcomputepricing.Standard)):
			opts.DiskType = gcpcomputepricing.Standard
		case strings.Contains(strings.ToLower(*dt.Name), strings.ToLower(gcpcomputepricing.SSD)):
			opts.DiskType = gcpcomputepricing.SSD
		case strings.Contains(strings.ToLower(*dt.Name), strings.ToLower(gcpcomputepricing.Balanced)):
			opts.DiskType = gcpcomputepricing.Balanced
		}
	}
	if opts.DiskType != "" {
		monthlyRate, err := gcpcomputepricing.CalculateDisk(data.pricing, opts)
		if err != nil {
			return 0, err
		}
		startOfMonth := utils.BeginningOfMonth(time.Now())
		endOfMonth := utils.EndingOfMonth(time.Now())

		hourlyRate := monthlyRate / uint64(endOfMonth.Sub(startOfMonth).Hours())
		hourlyUsage := int(time.Since(startOfMonth).Hours())

		normalizedHourlyRate := float64(hourlyRate) / 1000000000
		cost = normalizedHourlyRate * float64(hourlyUsage)
	}

	return cost, nil
}
