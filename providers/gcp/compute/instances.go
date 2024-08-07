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
			if strings.Contains(err.Error(), "SERVICE_DISABLED") {
				logrus.Warn(err.Error())
				return resources, nil
			} else {
				logrus.WithError(err).Errorf("failed to list instances")
				return resources, err
			}
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

			cost, err := gcpcomputepricing.CalculateMachineCost(ctx, client, gcpcomputepricing.CalculateMachineCostData{
				MachineType:       instance.GetMachineType(),
				Project:           client.GCPClient.Credentials.ProjectID,
				Zone:              zone,
				Commitment:        resolveCommitment(instance),
				CreationTimestamp: instance.GetCreationTimestamp(),
				Pricing:           actualPricing,
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

// resolveCommitment resolve whether the instance is preemptible or on-demand.
func resolveCommitment(instance *computepb.Instance) string {
	if instance.Scheduling.Preemptible != nil && *instance.Scheduling.Preemptible {
		return gcpcomputepricing.Spot
	}
	return gcpcomputepricing.OnDemand
}
