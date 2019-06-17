package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type SSHKey struct {
	ID string `json:"id"`
}

func (ovh OVH) GetSSHKeys() (int, error) {
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
		keys := make([]SSHKey, 0)
		err = client.Get(fmt.Sprintf("/cloud/project/%s/sshkey", project.ID), &keys)
		if err != nil {
			return total, err
		}

		total += len(keys)
	}

	return total, nil
}
