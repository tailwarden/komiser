package container

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils"
	"google.golang.org/api/option"

	container "cloud.google.com/go/container/apiv1"
	containerpb "cloud.google.com/go/container/apiv1/containerpb"
)

func Clusters(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	clusterClient, err := container.NewClusterManagerClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create cluster client")
		return resources, err
	}

	req := &containerpb.ListClustersRequest{
		ProjectId: client.GCPClient.Credentials.ProjectID,
	}
	clusters, err := clusterClient.ListClusters(ctx, req)
	if err != nil {
		logrus.WithError(err).Errorf("failed to collect clusters")
		return resources, err
	}

	for _, cluster := range clusters.Clusters {
		tags := make([]models.Tag, 0)
		//according to docs, NodeConfig is deprecated and is advised to use node_pool.config
		if cluster.NodeConfig.Labels != nil {
			for key, value := range cluster.NodeConfig.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}
		}
		zone := utils.GcpExtractZoneFromURL(cluster.GetLocation())

		resources = append(resources, models.Resource{
			Provider:   "GCP",
			Account:    client.Name,
			Service:    "Cluster",
			ResourceId: cluster.GetId(),
			Region:     zone,
			Name:       cluster.GetName(),
			FetchedAt:  time.Now(),
			Tags:       tags,
			Link:       fmt.Sprintf("https://console.cloud.google.com/kubernetes/clusters/details/%s/%s/details?project=%s", zone, cluster.GetName(), client.GCPClient.Credentials.ProjectID),
		})
	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "Container Clusters",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
