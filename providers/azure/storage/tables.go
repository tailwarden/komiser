package storage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Tables(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	storage := make([]Storage, 0)

	storageAccountsClient, err := armstorage.NewAccountsClient(client.AzureClient.SubscriptionId, client.AzureClient.Credentials, &arm.ClientOptions{})
	if err != nil {
		return resources, err
	}

	storagePager := storageAccountsClient.NewListPager(nil)
	for storagePager.More() {
		page, err := storagePager.NextPage(ctx)
		if err != nil {
			return resources, err
		}

		for _, v := range page.Value {
			storage = append(storage, Storage{
				Name:          *v.Name,
				ResourceGroup: strings.Split(*v.ID, "/")[4],
				Location:      *v.Location,
			})

		}
	}

	svc, err := armstorage.NewTableClient(
		client.AzureClient.SubscriptionId,
		client.AzureClient.Credentials,
		&arm.ClientOptions{},
	)
	if err != nil {
		return resources, err
	}

	for _, v := range storage {
		pager := svc.NewListPager(v.ResourceGroup, v.Name, &armstorage.TableClientListOptions{})
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				return resources, err
			}

			for _, table := range page.Value {

				tags := make([]models.Tag, 0)


				resources = append(resources, models.Resource{
					Provider:   "Azure",
					Account:    client.Name,
					Service:    "Queue",
					Region:     v.Location,
					ResourceId: *table.ID,
					Cost:       0,
					Name:       *table.Name,
					FetchedAt:  time.Now(),
					Tags:       tags,
					Link:       fmt.Sprintf("https://portal.azure.com/#resource%s", *table.ID),
				})
			}
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Azure",
		"account":   client.Name,
		"service":   "Table",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
