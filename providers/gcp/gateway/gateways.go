package gateway

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/apigateway/v1"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func ApiGateways(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	regions, err := listGCPRegions(client.GCPClient.Credentials.ProjectID, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to list zones to fetch api gateways")
		return resources, err
	}

	apiGatewayService, err := apigateway.NewService(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create API Gateway service")
		return resources, err
	}

RegionsLoop:
	for _, regionName := range regions {
		apiGateways, err := apiGatewayService.Projects.Locations.Gateways.List(
			"projects/" + client.GCPClient.Credentials.ProjectID + "/locations/" + regionName,
		).Do()
		if err != nil {
			if err.Error() == "googleapi: Error 403: Location "+regionName+" is not found or access is unauthorized., forbidden" || strings.Contains(err.Error(), "SERVICE_DISABLED") {
				logrus.Warn(err.Error())
				continue RegionsLoop
			} else {
				logrus.WithError(err).Errorf("failed to list API Gateways")
				return resources, err
			}
		}

		for _, apiGateway := range apiGateways.Gateways {
			parsedCreatedTime, err := time.Parse(time.RFC3339Nano, apiGateway.CreateTime)
			if err != nil {
				logrus.WithError(err).Errorf("failed to parse create time for API Gateways")
				return resources, err
			}

			resources = append(resources, models.Resource{
				Provider:   "GCP",
				Account:    client.Name,
				Service:    "API Gateways",
				ResourceId: apiGateway.Name,
				Name:       apiGateway.DisplayName,
				CreatedAt:  parsedCreatedTime,
				Region:     regionName,
				Metadata: map[string]string{
					"API Config":       apiGateway.ApiConfig,
					"Default Hostname": apiGateway.DefaultHostname,
					"State":            apiGateway.State,
				},
				FetchedAt: time.Now(),
				Link:      fmt.Sprintf("https://console.cloud.google.com/api-gateway/gateway/%s/location/%s?project=%s", apiGateway.DisplayName, regionName, client.GCPClient.Credentials.ProjectID),
			})

		}

	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "API Gateway",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}

func listGCPRegions(projectId string, creds option.ClientOption) ([]string, error) {
	var regions []string

	ctx := context.Background()
	computeService, err := compute.NewService(ctx, creds)
	if err != nil {
		logrus.WithError(err).Errorf("failed to create new service for fetching GCP regions for api gateway")
		return nil, err
	}

	regionList, err := computeService.Regions.List(projectId).Do()
	if err != nil {
		logrus.WithError(err).Errorf("failed to list regions for fetching GCP regions for api gateway")
		return nil, err
	}

	for _, region := range regionList.Items {
		regions = append(regions, region.Name)
	}
	return regions, nil
}
