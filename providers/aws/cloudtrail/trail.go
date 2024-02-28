package cloudtrail

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Trails(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config cloudtrail.ListTrailsInput
	resources := make([]models.Resource, 0)
	neptuneClient := cloudtrail.NewFromConfig(*client.AWSClient)

	output, err := neptuneClient.ListTrails(ctx, &config)
	if err != nil {
		return resources, err
	}

	for _, trail := range output.Trails {
		trailName := ""
		if trail.Name != nil {
			trailName = *trail.Name
		}
		resources = append(resources, models.Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "Cloudtrail Trail",
			Region:     client.AWSClient.Region,
			ResourceId: *trail.TrailARN,
			Name:       trailName,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/cloudtrailv2/home?region=%s#/trails/%s/%s", client.AWSClient.Region, client.AWSClient.Region, *trail.TrailARN, trailName),
		})
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Cloudtrail Trail",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
