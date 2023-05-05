package compute

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/PumpkinSeed/gcpcomputepricing"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils"
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
		logrus.WithError(err).Errorf("failedto fetch actual GCP VM pricing") // TODO figure out text
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

			cost, err := calculateCost(ctx, client, instance.GetMachineType(), client.GCPClient.Credentials.ProjectID, zone, actualPricing)
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

func calculateCost(ctx context.Context, client providers.ProviderClient, machineType, project, zone string, pricing *gcpcomputepricing.Pricing) (float64, error) {
	machineTypeClient, err := compute.NewMachineTypesRESTClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		return 0, err
	}

	mtS := strings.Split(machineType, "/")

	mt, err := machineTypeClient.Get(ctx, &computepb.GetMachineTypeRequest{
		MachineType: mtS[len(mtS)-1],
		Project:     project,
		Zone:        zone,
	})
	if err != nil {
		return 0, err
	}

	var cost float64
	if mt.Name != nil {
		switch {
		case strings.Contains(*mt.Name, gcpcomputepricing.E2):
			hourlyRate, err := gcpcomputepricing.Calculate(pricing, gcpcomputepricing.Opts{
				Type:        gcpcomputepricing.E2,
				Commitment:  gcpcomputepricing.OnDemand, // TODO
				Region:      "us-central1",
				NumOfCPU:    uint64(*mt.GuestCpus),
				NumOfMemory: uint64(*mt.MemoryMb / 1024),
			})
			if err != nil {
				return 0, err
			}
			cost = float64(hourlyRate) * 24 * 30 // TODO
		}
	}

	fmt.Println(mt)

	return cost / 1000000000, nil
}
