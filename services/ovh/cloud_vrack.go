package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type CloudIp struct {
	ID int `json:"id"`
}

type Network struct {
	ID string `json:"id"`
}

type vRack struct {
	ID string `json:"id"`
}

func (ovh OVH) GetIps() (int, error) {
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
		ips := make([]CloudIp, 0)
		err = client.Get(fmt.Sprintf("/cloud/project/%s/ip", project.ID), &ips)
		if err != nil {
			return total, err
		}

		total += len(ips)
	}

	return total, nil
}

func (ovh OVH) GetPrivateNetworks() (int, error) {
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
		networks := make([]Network, 0)
		err = client.Get(fmt.Sprintf("/cloud/project/%s/network/private", project.ID), &networks)
		if err != nil {
			return total, err
		}

		total += len(networks)
	}

	return total, nil
}

func (ovh OVH) GetPublicNetworks() (int, error) {
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
		networks := make([]Network, 0)
		err = client.Get(fmt.Sprintf("/cloud/project/%s/network/public", project.ID), &networks)
		if err != nil {
			return total, err
		}

		total += len(networks)
	}

	return total, nil
}

func (ovh OVH) GetFailoverIps() (int, error) {
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
		ips := make([]Network, 0)
		err = client.Get(fmt.Sprintf("/cloud/project/%s/ip/failover", project.ID), &ips)
		if err != nil {
			return total, err
		}

		total += len(ips)
	}

	return total, nil
}

func (ovh OVH) GetVRacks() (int, error) {
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
		vrack := vRack{}
		err = client.Get(fmt.Sprintf("/cloud/project/%s/vrack", project.ID), &vrack)
		if err != nil {
			return total, err
		}

		if len(vrack.ID) != 0 {
			total++
		}
	}

	return total, nil
}
