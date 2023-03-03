package containers

import (
	"context"
	"fmt"
	"github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"time"

	"github.com/scaleway/scaleway-sdk-go/scw"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func K8sClusters(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	k8sSvc := k8s.NewAPI(client.ScalewayClient)

	regions := []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}

	for _, region := range regions {
		output, err := k8sSvc.ListClusters(&k8s.ListClustersRequest{
			Region: region,
		})
		if err != nil {
			return resources, err
		}

		for _, k8sCluster := range output.Clusters {
			resources = append(resources, models.Resource{
				Provider:   "Scaleway",
				Account:    client.Name,
				Service:    "Kubernetes",
				Region:     k8sCluster.Region.String(),
				ResourceId: k8sCluster.ID,
				Cost:       0,
				Name:       k8sCluster.Name,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://console.scaleway.com/kapsule/clusters/%s/%s", k8sCluster.Region.String(), k8sCluster.ID),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Scaleway",
		"account":   client.Name,
		"service":   "Kubernetes",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
