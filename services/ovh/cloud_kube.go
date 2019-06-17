package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type Node struct {
	ID string `json:"id"`
}

func (ovh OVH) GetKubeClusters() (int, error) {
	total := 0

	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return total, err
	}

	projects, err := ovh.GetProjects()
	if err != nil {
		return total, err
	}

	for _, project := range projects {
		clusters := []string{}
		err = client.Get(fmt.Sprintf("/cloud/project/%s/kube", project.ID), &clusters)
		if err != nil {
			return total, err
		}

		total += len(clusters)
	}

	return total, nil
}

func (ovh OVH) GetKubeNodes() (int, error) {
	total := 0

	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return total, err
	}

	projects, err := ovh.GetProjects()
	if err != nil {
		return total, err
	}

	for _, project := range projects {
		clusters := []string{}
		err = client.Get(fmt.Sprintf("/cloud/project/%s/kube", project.ID), &clusters)
		if err != nil {
			return total, err
		}

		for _, clusterId := range clusters {
			nodes := make([]Node, 0)
			err = client.Get(fmt.Sprintf("/cloud/project/%s/kube/%s/node", project.ID, clusterId), &nodes)
			if err != nil {
				return total, err
			}

			total += len(nodes)
		}
	}

	return total, nil
}
