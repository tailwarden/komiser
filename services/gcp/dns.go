package gcp

import (
	"context"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	dns "google.golang.org/api/dns/v1"
)

func (gcp GCP) GetManagedZones() (int, error) {
	sum := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, dns.CloudPlatformReadOnlyScope)
	if err != nil {
		return sum, err
	}

	client := oauth2.NewClient(context.Background(), src)

	svc, err := dns.New(client)
	if err != nil {
		return sum, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return sum, err
	}

	for _, project := range projects {
		zones, err := svc.ManagedZones.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return 0, err
		}
		sum += len(zones.ManagedZones)
	}

	return sum, nil
}

func (gcp GCP) GetARecords() (int, error) {
	sum := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, dns.CloudPlatformReadOnlyScope)
	if err != nil {
		return sum, err
	}

	client := oauth2.NewClient(context.Background(), src)

	svc, err := dns.New(client)
	if err != nil {
		return sum, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return sum, err
	}

	for _, project := range projects {
		zones, err := svc.ManagedZones.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return sum, err
		}

		for _, zone := range zones.ManagedZones {
			records, err := svc.ResourceRecordSets.List(project.ID, zone.Name).Do()
			if err != nil {
				log.Println(err)
				return sum, err
			}

			for _, record := range records.Rrsets {
				if record.Type == "A" {
					sum++
				}
			}
		}
	}

	return sum, nil
}
