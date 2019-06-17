package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type Volume struct {
	Type       string   `json:"type"`
	Region     string   `json:"region"`
	Size       int      `json:"size"`
	AttachedTo []string `json:"attachedTo"`
}

func (ovh OVH) GetVolumes() ([]Volume, error) {
	listOfVolumes := make([]Volume, 0)

	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return listOfVolumes, err
	}

	projects, err := ovh.GetProjects()
	if err != nil {
		return listOfVolumes, err
	}

	for _, project := range projects {
		volumes := make([]Volume, 0)
		err = client.Get(fmt.Sprintf("/cloud/project/%s/volume", project.ID), &volumes)
		if err != nil {
			return listOfVolumes, err
		}

		for _, volume := range volumes {
			listOfVolumes = append(listOfVolumes, volume)
		}
	}

	return listOfVolumes, nil
}
