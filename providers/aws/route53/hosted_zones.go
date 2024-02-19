package route53

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func HostedZones(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config route53.ListHostedZonesInput
	resources := make([]models.Resource, 0)
	route53Client := route53.NewFromConfig(*client.AWSClient)

	output, err := route53Client.ListHostedZones(ctx, &config)
	if err != nil {
		return resources, err
	}

	for _, zone := range output.HostedZones {
		resources = append(resources, models.Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "Route53 Hosted Zone",
			Region:     client.AWSClient.Region,
			ResourceId: *zone.Id,
			Name:       *zone.Name,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/route53/v2/hostedzones#ListRecordSets/%s", client.AWSClient.Region, *zone.Id),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Route53 Hosted Zone",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
