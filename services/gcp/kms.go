package gcp

import (
	"context"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	cloudkms "google.golang.org/api/cloudkms/v1"
)

func (gcp GCP) GetKMSCryptoKeys() (int, error) {
	total := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, cloudkms.CloudPlatformScope)
	if err != nil {
		return total, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := cloudkms.New(client)
	if err != nil {
		return total, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return total, err
	}

	for _, project := range projects {
		locations, err := svc.Projects.Locations.List("projects/" + project.ID).Do()
		if err != nil {
			log.Println(err)
			return total, err
		}
		for _, location := range locations.Locations {
			keyRings, err := svc.Projects.Locations.KeyRings.List("projects/" + project.ID + "/locations/" + location.LocationId).Do()
			if err != nil {
				log.Println(err)
				return total, err
			}
			for _, keyRing := range keyRings.KeyRings {
				keys, err := svc.Projects.Locations.KeyRings.CryptoKeys.List("projects/" + project.ID + "/locations/" + location.LocationId + "/keyRings/" + keyRing.Name).Do()
				if err != nil {
					log.Println(err)
					return total, err
				}
				total += len(keys.CryptoKeys)
			}
		}
	}
	return total, nil
}
