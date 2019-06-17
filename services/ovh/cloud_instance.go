package ovh

import (
	"fmt"
	"strings"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type Instance struct {
	ID          string `json:"id"`
	Region      string `json:"region"`
	Status      string `json:"status"`
	ImageId     string `json:"imageId"`
	Model       string `json:"planCode"`
	FlavorId    string `json:"flavorId"`
	IPAddresses []struct {
		GatewayIp string `json:"gatewayIp"`
		Type      string `json:"type"`
	} `json:"ipAddresses"`
}

func (ovh OVH) GetInstances() ([]Instance, error) {
	listOfInstances := make([]Instance, 0)

	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return listOfInstances, err
	}

	projects, err := ovh.GetProjects()
	if err != nil {
		return listOfInstances, err
	}

	for _, project := range projects {
		instances := make([]Instance, 0)
		err = client.Get(fmt.Sprintf("/cloud/project/%s/instance", project.ID), &instances)
		if err != nil {
			return listOfInstances, err
		}

		for _, instance := range instances {
			instance.Model = strings.Replace(instance.Model, ".consumption", "", -1)
			listOfInstances = append(listOfInstances, instance)
		}
	}

	return listOfInstances, nil
}
