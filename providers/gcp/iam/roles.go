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

func Roles(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	iamService, err := iam.NewService(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create IAM roles service")
		return resources, err
	}

	roles, err := iamService.Projects.Roles.List(
		"projects/" + client.GCPClient.Credentials.ProjectID,
	).Do()
	if err != nil {
		if strings.Contains(err.Error(), "SERVICE_DISABLED") {
			logrus.Warn(err.Error())
			return resources, nil
		} else {
			logrus.WithError(err).Errorf("failed to list IAM roles")
			return resources, err
		}
	}

	for _, role := range roles.Roles {
		targetForUrl := strings.Replace(role.Name, "/", "<", -1)

		resources = append(resources, models.Resource{
			Provider:   "GCP",
			Account:    client.Name,
			Service:    "IAM Roles",
			ResourceId: role.Name,
			Name:       role.Title,
			Metadata: map[string]string{
				"Description": role.Description,
			},
			FetchedAt: time.Now(),
			Link:      fmt.Sprintf("https://console.cloud.google.com/iam-admin/roles/details/%s?project=%s", targetForUrl, client.GCPClient.Credentials.ProjectID),
		})
	}
	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "IAM Custom Roles",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil

}
