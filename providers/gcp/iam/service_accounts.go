package iam

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
)

func ServiceAccounts(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	iamService, err := iam.NewService(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create IAM Service Account service")
		return resources, err
	}

	serviceAccounts, err := iamService.Projects.ServiceAccounts.List(
		"projects/" + client.GCPClient.Credentials.ProjectID,
	).Do()
	if err != nil {
		if strings.Contains(err.Error(), "SERVICE_DISABLED") {
			logrus.Warn(err.Error())
			return resources, nil
		} else {
			logrus.WithError(err).Errorf("failed to list IAM Service Accounts")
			return resources, err
		}
	}

	for _, account := range serviceAccounts.Accounts {
		resources = append(resources, models.Resource{
			Provider:   "GCP",
			Account:    client.Name,
			Service:    "IAM Service Accounts",
			ResourceId: account.UniqueId,
			Name:       account.DisplayName,
			Metadata: map[string]string{
				"Description": account.Description,
				"Email":       account.Email,
			},
			FetchedAt: time.Now(),
			Link:      fmt.Sprintf("https://console.cloud.google.com/iam-admin/serviceaccounts/details/%s?project=%s", account.UniqueId, client.GCPClient.Credentials.ProjectID),
		})

	}
	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "IAM Service Accounts",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil

}
