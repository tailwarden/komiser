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
	computepb "cloud.google.com/go/compute/apiv1/computepb"
)

func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	instancesClient, err := compute.NewInstancesRESTClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create compute client")
		return resources, err
	}

	req := &computepb.AggregatedListInstancesRequest{
		Project: client.GCPClient.Credentials.ProjectID,
	}
	instances := instancesClient.AggregatedList(ctx, req)

	actualPricing, err := gcpcomputepricing.Fetch()
	if err != nil {
		logrus.WithError(err).Errorf("failed to fetch actual GCP VM pricing")
		return resources, err
	}

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

			zone := utils.GcpExtractZoneFromURL(instance.GetZone())

			cost, err := calculateCost(ctx, client, calculateCostData{
				machineType: instance.GetMachineType(),
				project:     client.GCPClient.Credentials.ProjectID,
				zone:        zone,
				commitment:  resolveCommitment(instance),
				pricing:     actualPricing,
			})
			if err != nil {
				logrus.WithError(err).Errorf("failed to calculate cost")
			}

			resources = append(resources, models.Resource{
				Provider:   "GCP",
				Account:    client.Name,
				Service:    "VM Instance",
				ResourceId: fmt.Sprintf("%d", instance.GetId()),
				Region:     zone,
				Name:       instance.GetName(),
				Cost:       cost,
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

type calculateCostData struct {
	machineType string
	project     string
	zone        string
	commitment  string
	pricing     *gcpcomputepricing.Pricing
}

func calculateCost(ctx context.Context, client providers.ProviderClient, data calculateCostData) (float64, error) {
	machineTypeClient, err := compute.NewMachineTypesRESTClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		return 0, err
	}

	mtS := strings.Split(data.machineType, "/")

	mt, err := machineTypeClient.Get(ctx, &computepb.GetMachineTypeRequest{
		MachineType: mtS[len(mtS)-1],
		Project:     data.project,
		Zone:        data.zone,
	})
	if err != nil {
		return 0, err
	}

	var opts = gcpcomputepricing.Opts{
		Commitment:  data.commitment,
		Region:      utils.GcpGetRegionFromZone(data.zone),
		NumOfCPU:    uint64(*mt.GuestCpus),
		NumOfMemory: uint64(*mt.MemoryMb / 1024),
	}
	var cost float64
	if mt.Name != nil {
		switch {
		case strings.Contains(strings.ToLower(*mt.Name), strings.ToLower(gcpcomputepricing.E2)):
			opts.Type = gcpcomputepricing.E2
		case strings.Contains(strings.ToLower(*mt.Name), strings.ToLower(gcpcomputepricing.C3)):
			opts.Type = gcpcomputepricing.C3
		case strings.Contains(strings.ToLower(*mt.Name), strings.ToLower(gcpcomputepricing.N2)):
			opts.Type = gcpcomputepricing.N2
		case strings.Contains(strings.ToLower(*mt.Name), strings.ToLower(gcpcomputepricing.N2D)):
			opts.Type = gcpcomputepricing.N2D
		case strings.Contains(strings.ToLower(*mt.Name), strings.ToLower(gcpcomputepricing.T2A)):
			opts.Type = gcpcomputepricing.T2A
		case strings.Contains(strings.ToLower(*mt.Name), strings.ToLower(gcpcomputepricing.T2D)):
			opts.Type = gcpcomputepricing.T2D
		case strings.Contains(strings.ToLower(*mt.Name), strings.ToLower(gcpcomputepricing.N1)):
			opts.Type = gcpcomputepricing.N1
		case strings.Contains(strings.ToLower(*mt.Name), strings.ToLower(gcpcomputepricing.C2)):
			opts.Type = gcpcomputepricing.C2
		case strings.Contains(strings.ToLower(*mt.Name), strings.ToLower(gcpcomputepricing.C2D)):
			opts.Type = gcpcomputepricing.C2D
		case strings.Contains(strings.ToLower(*mt.Name), strings.ToLower(gcpcomputepricing.M1)):
			opts.Type = gcpcomputepricing.M1
		case strings.Contains(strings.ToLower(*mt.Name), strings.ToLower(gcpcomputepricing.M2)):
			opts.Type = gcpcomputepricing.M2
		case strings.Contains(strings.ToLower(*mt.Name), strings.ToLower(gcpcomputepricing.M3)):
			opts.Type = gcpcomputepricing.M3
		}
	}
	if opts.Type != "" {
		hourlyRate, err := gcpcomputepricing.CalculateMachine(data.pricing, opts)
		if err != nil {
			return 0, err
		}
		startOfMonth := utils.BeginningOfMonth(time.Now())
		hourlyUsage := int(time.Since(startOfMonth).Hours())
		normalizedHourlyRate := float64(hourlyRate) / 1000000000
		cost = normalizedHourlyRate * float64(hourlyUsage)
	}

	return cost, nil
}

func resolveCommitment(instance *computepb.Instance) string {
	if instance.Scheduling.Preemptible != nil && *instance.Scheduling.Preemptible {
		return gcpcomputepricing.Spot
	}
	return gcpcomputepricing.OnDemand
}
