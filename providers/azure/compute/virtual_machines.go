package compute

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func VirtualMachines(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	svc, err := armcompute.NewVirtualMachinesClient(client.AzureClient.SubscriptionId, client.AzureClient.Credentials, &arm.ClientOptions{})
	if err != nil {
		return resources, err
	}

	costClient, err := armcostmanagement.NewQueryClient(client.AzureClient.Credentials, &policy.ClientOptions{})
	if err != nil {
		return resources, err
	}

	pager := svc.NewListAllPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return resources, err
		}

		for _, vm := range page.VirtualMachineListResult.Value {
			tags := make([]models.Tag, 0)
			queryResult, err := costClient.Usage(ctx, "subscriptions/"+client.AzureClient.SubscriptionId, armcostmanagement.QueryDefinition{
				Type: to.Ptr(armcostmanagement.ExportTypeUsage),
				Dataset: &armcostmanagement.QueryDataset{
					Aggregation: map[string]*armcostmanagement.QueryAggregation{
						"totalCost": {
							Name:     to.Ptr("PreTaxCost"),
							Function: to.Ptr(armcostmanagement.FunctionTypeSum),
						},
					},
					Granularity: to.Ptr(armcostmanagement.GranularityType("None")),
				},
				Timeframe: to.Ptr(armcostmanagement.TimeframeTypeMonthToDate),
			}, nil)
			if err != nil {
				log.Warnf("failed to query usage: %v\n", err)
			}

			cost := queryResult.Properties.Rows[0][0].(float64)

			for key, value := range vm.Tags {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: *value,
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "Azure",
				Account:    client.Name,
				Service:    "Virtual Machine",
				Region:     *vm.Location,
				ResourceId: *vm.ID,
				Cost:       cost,
				Name:       *vm.Name,
				FetchedAt:  time.Now(),
				Relations:  getVmRelation(vm),
				Tags:       tags,
				Link:       fmt.Sprintf("https://portal.azure.com/#resource%s", *vm.ID),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Azure",
		"account":   client.Name,
		"service":   "Virtual Machine",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func getVmRelation(vm *armcompute.VirtualMachine) []models.Link {

	var rel []models.Link

	if vm.Properties.StorageProfile != nil && vm.Properties.StorageProfile.DataDisks != nil {
		for _, disk := range vm.Properties.StorageProfile.DataDisks {
			rel = append(rel, models.Link{
				ResourceID: *disk.ManagedDisk.ID,
				Type:       "Disk",
				Name:       *disk.Name,
				Relation:   "USES",
			})
		}

		if vm.Properties.StorageProfile.ImageReference != nil {
			rel = append(rel, models.Link{
				ResourceID: *vm.Properties.StorageProfile.ImageReference.ID,
				Type:       "Image",
				Name:       *vm.Properties.StorageProfile.ImageReference.ID,
				Relation:   "USES",
			})
		}
	}
	return rel
}
