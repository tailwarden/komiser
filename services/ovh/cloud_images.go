package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type Image struct {
	Type string `json:"type"`
}

type Images struct {
	Windows int `json:"windows"`
	Linux   int `json:"linux"`
}

func (ovh OVH) GetImages() (Images, error) {
	listOfImages := Images{}

	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return listOfImages, err
	}

	projects, err := ovh.GetProjects()
	if err != nil {
		return listOfImages, err
	}

	for _, project := range projects {
		images := make([]Image, 0)
		err = client.Get(fmt.Sprintf("/cloud/project/%s/image", project.ID), &images)
		if err != nil {
			return listOfImages, err
		}

		for _, img := range images {
			if img.Type == "windows" {
				listOfImages.Windows++
			} else {
				listOfImages.Linux++
			}
		}
	}

	return listOfImages, nil
}
