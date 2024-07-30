package alloydb

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"

	alloydb "cloud.google.com/go/alloydb/apiv1"
	"cloud.google.com/go/alloydb/apiv1/alloydbpb"
	"google.golang.org/api/iterator"
)

func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	alloydbClient, err := alloydb.NewAlloyDBAdminClient(ctx)
	if err != nil {
		logrus.WithError(err).Errorf("failed to create alloydb client")
		return resources, err
	}
	listInstancesParams := &alloydbpb.ListInstancesRequest{}
	instanceIterator := alloydbClient.ListInstances(ctx, listInstancesParams)
	for {
		instance, err := instanceIterator.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			if strings.Contains(err.Error(), "SERVICE_DISABLED") {
				logrus.Warn(err.Error())
				return resources, nil
			} else {
				logrus.WithError(err).Errorf("failed to get instances")
				return resources, err
			}
		}
		resources = append(resources, models.Resource{
			Provider:   "GCP",
			Account:    client.Name,
			Service:    "AlloyDB Instances",
			ResourceId: instance.Uid,
			Name:       instance.Name,
			CreatedAt:  instance.CreateTime.AsTime(),
			Cost:       0,
			FetchedAt:  time.Now(),
		})
	}
	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "AlloyDB Instances",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
