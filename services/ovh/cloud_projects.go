package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type Project struct {
	ID          string `json:"project_id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func (ovh OVH) GetProjects() ([]Project, error) {
	projects := make([]Project, 0)

	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return projects, err
	}

	ids := []string{}
	err = client.Get("/cloud/project", &ids)
	if err != nil {
		return projects, err
	}

	for _, id := range ids {
		project := Project{}
		err = client.Get(fmt.Sprintf("/cloud/project/%s", id), &project)
		if err != nil {
			return projects, err
		}

		projects = append(projects, project)
	}
	return projects, nil
}
