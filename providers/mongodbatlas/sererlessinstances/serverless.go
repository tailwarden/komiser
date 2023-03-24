package serverless

import (
	"context"
	"fmt"

	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils"
	"go.mongodb.org/atlas/mongodbatlas"
)

func ServerlessInstances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	projects, _, err := client.MongoDBAtlasClient.Projects.GetAllProjects(ctx, &mongodbatlas.ListOptions{})
	if err != nil {
		log.WithError(err).Error("Error getting projects")
	}

	for _, project := range projects.Results {
		clusters, _, err := client.MongoDBAtlasClient.ServerlessInstances.List(ctx, project.ID, &mongodbatlas.ListOptions{})
		if err != nil {
			log.WithError(err).Error("Error getting clusters")
		}

		for _, cluster := range clusters.Results {
			tags := make([]models.Tag, 0)
			for _, tag := range cluster.Labels {
				tags = append(tags, models.Tag{
					Key:   tag.Key,
					Value: tag.Value,
				})
			}

			monthlyCost := 0.0

			resources = append(resources, models.Resource{
				Provider:   "MongoDBAtlas",
				Account:    client.Name,
				Service:    "Serverless Cluster",
				Region:     utils.NormalizeRegionName(cluster.ProviderSettings.RegionName),
				ResourceId: cluster.ID,
				Name:       cluster.Name,
				Cost:       monthlyCost,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://cloud.mongodb.com/v2/%s#/clusters/detail/mvp", cluster.GroupID),
				Tags:       tags,
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "MongoDBAtlas",
		"account":   client.Name,
		"service":   "Serverless Cluster",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
