package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type Quota struct {
	ID      string `json:"region"`
	KeyPair struct {
		Max int `json:"maxCount"`
	} `json:"keypair"`
	Volume struct {
		MaxGigaBytes  int `json:"maxGigabytes"`
		UsedGigaBytes int `json:"usedGigabytes"`
		VolumeCount   int `json:"volumeCount"`
	} `json:"volume"`
	Instance struct {
		MaxCores      int `json:"maxCores"`
		MaxInstances  int `json:"maxInstances"`
		UsedCores     int `json:"usedCores"`
		UsedInstances int `json:"usedInstances"`
		MaxRAM        int `json:"maxRam"`
		UsedRAM       int `json:"usedRAM"`
	} `json:"instance"`
}

func (ovh OVH) GetLimits() ([]Quota, error) {
	limits := make([]Quota, 0)

	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return limits, err
	}

	ids := []string{}
	err = client.Get("/cloud/project", &ids)
	if err != nil {
		return limits, err
	}

	for _, id := range ids {
		quotas := make([]Quota, 0)
		err = client.Get(fmt.Sprintf("/cloud/project/%s/quota", id), &quotas)
		if err != nil {
			return limits, err
		}

		for _, q := range quotas {
			limits = append(limits, q)
		}
	}
	return limits, nil
}
