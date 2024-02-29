package secretsmanager

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Secrets(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config secretsmanager.ListSecretsInput
	resources := make([]models.Resource, 0)
	neptuneClient := secretsmanager.NewFromConfig(*client.AWSClient)

	output, err := neptuneClient.ListSecrets(ctx, &config)
	if err != nil {
		return resources, err
	}

	for _, secret := range output.SecretList {
		secretName := ""
		if secret.Name != nil {
			secretName = *secret.Name
		}
		resources = append(resources, models.Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "Secret",
			Region:     client.AWSClient.Region,
			ResourceId: *secret.ARN,
			Name:       secretName,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/secretsmanager/secret?name=%s&region=%s", client.AWSClient.Region, secretName, client.AWSClient.Region),
		})
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Secret",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
