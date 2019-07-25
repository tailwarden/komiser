package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

type Certificate struct {
	Custom      int `json:"custom"`
	LetsEncrypt int `json:"letsEncrypt"`
}

func (dg DigitalOcean) DescribeCertificates(client *godo.Client) (Certificate, error) {
	certificate := Certificate{}

	certificates, _, err := client.Certificates.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return certificate, err
	}

	for _, c := range certificates {
		if c.Type == "custom" {
			certificate.Custom++
		} else {
			certificate.LetsEncrypt++
		}
	}

	return certificate, nil
}
