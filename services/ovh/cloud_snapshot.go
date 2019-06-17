package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type Snapshot struct {
	Region string `json:"region"`
	Size   int    `json:"size"`
}

func (ovh OVH) GetSnapshots() ([]Snapshot, error) {
	listOfSnapshots := make([]Snapshot, 0)

	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return listOfSnapshots, err
	}

	projects, err := ovh.GetProjects()
	if err != nil {
		return listOfSnapshots, err
	}

	for _, project := range projects {
		snpashots := make([]Snapshot, 0)
		err = client.Get(fmt.Sprintf("/cloud/project/%s/volume/snapshot", project.ID), &snpashots)
		if err != nil {
			return listOfSnapshots, err
		}

		for _, snpashot := range snpashots {
			listOfSnapshots = append(listOfSnapshots, snpashot)
		}
	}

	return listOfSnapshots, nil
}
