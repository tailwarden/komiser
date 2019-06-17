package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type StorageContainer struct {
	StoredObjects int    `json:"storedObjects"`
	StoredBytes   int    `json:"storedBytes"`
	Region        string `json:"region"`
}

func (ovh OVH) GetStorageContainers() ([]StorageContainer, error) {
	listOfContainers := make([]StorageContainer, 0)

	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return listOfContainers, err
	}

	projects, err := ovh.GetProjects()
	if err != nil {
		return listOfContainers, err
	}

	for _, project := range projects {
		containers := make([]StorageContainer, 0)
		err = client.Get(fmt.Sprintf("/cloud/project/%s/storage", project.ID), &containers)
		if err != nil {
			return listOfContainers, err
		}

		for _, container := range containers {
			listOfContainers = append(listOfContainers, container)
		}
	}

	return listOfContainers, nil
}
