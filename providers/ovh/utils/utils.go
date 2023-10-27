package utils

import "github.com/tailwarden/komiser/providers"

func GetProjects(client providers.ProviderClient) ([]string, error) {
	projectIds := []string{}
	err := client.OVHClient.Get("/v2/cloud/project", projectIds)
	if err != nil {
		return nil, err
	}
	return projectIds, nil
}
