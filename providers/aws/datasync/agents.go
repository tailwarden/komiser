package datasync

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/datasync"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Agents(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config datasync.ListAgentsInput
	resources := make([]models.Resource, 0)
	neptuneClient := datasync.NewFromConfig(*client.AWSClient)

	output, err := neptuneClient.ListAgents(ctx, &config)
	if err != nil {
		return resources, err
	}

	for _, agent := range output.Agents {
		agentName := ""
		if agent.Name != nil {
			agentName = *agent.Name
		}
		resources = append(resources, models.Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "DataSync Agent",
			Region:     client.AWSClient.Region,
			ResourceId: *agent.AgentArn,
			Name:       agentName,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/datasync/home?region=%s#/agents", client.AWSClient.Region, client.AWSClient.Region),
		})
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "DataSync Agent",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
