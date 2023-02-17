package storage

import (
	"context"
	"time"

	"github.com/oracle/oci-go-sdk/objectstorage"

	log "github.com/sirupsen/logrus"

	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Buckets(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	objectStorageClient, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(client.OciClient)
	if err != nil {
		return resources, err
	}

	tenancyOCID, err := client.OciClient.TenancyOCID()
	if err != nil {
		return resources, err
	}

	getNamespaceRequestConfig := objectstorage.GetNamespaceRequest{
		CompartmentId: &tenancyOCID,
	}

	getNamespaceOutput, err := objectStorageClient.GetNamespace(context.Background(), getNamespaceRequestConfig)
	if err != nil {
		return resources, err
	}

	listBucketsRequestConfig := objectstorage.ListBucketsRequest{
		CompartmentId: &tenancyOCID,
		NamespaceName: getNamespaceOutput.Value,
	}

	output, err := objectStorageClient.ListBuckets(context.Background(), listBucketsRequestConfig)
	if err != nil {
		return resources, err
	}

	for _, bucket := range output.Items {
		tags := make([]Tag, 0)

		for key, value := range bucket.FreeformTags {
			tags = append(tags, Tag{
				Key:   key,
				Value: value,
			})
		}

		region, err := client.OciClient.Region()
		if err != nil {
			return resources, err
		}

		resources = append(resources, Resource{
			Provider:   "OCI",
			Account:    client.Name,
			ResourceId: *bucket.Name,
			Service:    "ObjectStorage Bucket",
			Region:     region,
			Name:       *bucket.Name,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "OCI",
		"account":   client.Name,
		"service":   "ObjectStorage Bucket",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
