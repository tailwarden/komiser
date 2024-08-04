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

func Clusters(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	alloydbClient, err := alloydb.NewAlloyDBAdminClient(ctx)
	if err != nil {
		logrus.WithError(err).Errorf("failed to create alloydb client")
		return resources, err
	}
	listClustersParams := &alloydbpb.ListClustersRequest{}
	clusterIterator := alloydbClient.ListClusters(ctx, listClustersParams)
	for {
		cluster, err := clusterIterator.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			if strings.Contains(err.Error(), "SERVICE_DISABLED") {
				logrus.Warn(err.Error())
				return resources, nil
			} else {
				logrus.WithError(err).Errorf("failed to get clusters")
				return resources, err
			}
		}
		resources = append(resources, models.Resource{
			Provider:   "GCP",
			Account:    client.Name,
			Service:    "AlloyDB Clusters",
			ResourceId: cluster.Uid,
			Name:       cluster.Name,
			CreatedAt:  cluster.CreateTime.AsTime(),
			Cost:       0,
			FetchedAt:  time.Now(),
		})
	}
	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "AlloyDB Clusters",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
