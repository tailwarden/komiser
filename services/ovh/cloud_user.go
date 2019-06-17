package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type User struct {
	ID int `json:"id"`
}

func (ovh OVH) GetUsers() (int, error) {
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
		users := make([]User, 0)
		err = client.Get(fmt.Sprintf("/cloud/project/%s/user", project.ID), &users)
		if err != nil {
			return total, err
		}

		total += len(users)
	}

	return total, nil
}
