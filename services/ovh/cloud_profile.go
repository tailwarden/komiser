package ovh

import (
	ovhClient "github.com/ovh/go-ovh/ovh"
)

type Profile struct {
	ID           string `json:"id"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Country      string `json:"country"`
	CustomerCode string `json:"customerCode"`
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	Name         string `json:"name"`
	Nichandle    string `json:"nichandle"`
	Zip          string `json:"zip"`
	Currency     struct {
		Code string `json:"code"`
	} `json:"currency"`
}

func (ovh OVH) GetProfile() (Profile, error) {
	profile := Profile{}

	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return profile, err
	}

	err = client.Get("/me", &profile)
	return profile, err
}
