package clusters

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

func getPricingMap(instanceSize string) float64 {
	// Values taken from https://www.mongodb.com/pricing
	pricingMap := map[string]float64{
		"M10":  0.08,
		"M20":  0.2,
		"M30":  0.54,
		"M40":  1.04,
		"M50":  2.0,
		"M60":  3.95,
		"M80":  7.30,
		"M140": 10.99,
		"M200": 14.59,
		"M300": 21.85,
		"M400": 22.40,
		"M700": 33.26,
	}

	return pricingMap[instanceSize]
}


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
			tags := make([]models.Tag, 0)
			for _, tag := range cluster.Labels {
				tags = append(tags, models.Tag{
					Key:   tag.Key,
					Value: tag.Value,
				})
			}

			monthlyCost := 0.0
			// M0 is the "shared" tier, which is free
			if cluster.ProviderSettings.InstanceSizeName != "M0" {
				startOfMonth := utils.BeginningOfMonth(time.Now())
				hourlyUsage := 0

				clusterCreatedAt, err := time.Parse(time.RFC3339, cluster.CreateDate)
				if err != nil {
					log.WithFields(log.Fields{
						"createDate": cluster.CreateDate,
					}).WithError(err).Errorf("Could not parse cluster.CreateDate")
				}

				if clusterCreatedAt.Before(startOfMonth) {
					hourlyUsage = int(time.Since(startOfMonth).Hours())
				} else {
					hourlyUsage = int(time.Since(clusterCreatedAt).Hours())
				}

				monthlyCost = float64(hourlyUsage) * getPricingMap(cluster.ProviderSettings.InstanceSizeName)
			}

			resources = append(resources, models.Resource{
				Provider:   "MongoDBAtlas",
				Account:    client.Name,
				Service:    "Cluster",
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
		"service":   "Cluster",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
