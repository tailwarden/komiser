package clusters

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"go.mongodb.org/atlas/mongodbatlas"
)

func Clusters(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	projects, _, err := client.MongoDBAtlasClient.Projects.GetAllProjects(ctx, &mongodbatlas.ListOptions{})
	if err != nil {
		log.WithError(err).Error("Error getting projects")
	}

	for _, project := range projects.Results {
		clusters, _, err := client.MongoDBAtlasClient.Clusters.List(ctx, project.ID, &mongodbatlas.ListOptions{})
		if err != nil {
			log.WithError(err).Error("Error getting clusters")
		}

		for _, cluster := range clusters {
			resources = append(resources, models.Resource{
				Provider:   "MongoDBAtlas",
				Account:    client.Name,
				Service:    "Cluster",
				Region:     cluster.ProviderSettings.RegionName,
				ResourceId: cluster.ID,
				Name:       cluster.Name,
				Cost:       0,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://cloud.mongodb.com/v2/%s#/clusters/detail/mvp", cluster.GroupID),
			})
		}

	}

	log.WithFields(log.Fields{
		"provider":  "MongoDBAtlas",
		"account":   client.Name,
		"service":   "Cluster",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
