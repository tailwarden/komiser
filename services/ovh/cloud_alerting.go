package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

func (ovh OVH) GetAlerts() (int, error) {
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
		alerts := []string{}
		err = client.Get(fmt.Sprintf("/cloud/project/%s/alerting", project.ID), &alerts)
		if err != nil {
			return total, err
		}

		total += len(alerts)
	}

	return total, nil
}
