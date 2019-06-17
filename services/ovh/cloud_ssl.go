package ovh

import (
	ovhClient "github.com/ovh/go-ovh/ovh"
)

func (ovh OVH) GetSSLCertificates() (int, error) {
	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return 0, err
	}

	certificates := []string{}
	err = client.Get("/ssl", &certificates)
	if err != nil {
		return 0, err
	}

	return len(certificates), nil
}

func (ovh OVH) GetSSLGateways() (int, error) {
	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return 0, err
	}

	certificates := []string{}
	err = client.Get("/sslGateway", &certificates)
	if err != nil {
		return 0, err
	}

	return len(certificates), nil
}
