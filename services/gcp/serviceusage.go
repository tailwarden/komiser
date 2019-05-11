package gcp

import (
	"context"
	"log"

	. "github.com/mlabouardy/komiser/models/gcp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	serviceusage "google.golang.org/api/serviceusage/v1"
)

func (gcp GCP) GetEnabledAPIs() ([]API, error) {
	listOfAPIs := make([]API, 0)

	src, err := google.DefaultTokenSource(oauth2.NoContext, serviceusage.CloudPlatformReadOnlyScope)
	if err != nil {
		return listOfAPIs, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := serviceusage.New(client)
	if err != nil {
		return listOfAPIs, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return listOfAPIs, err
	}

	for _, project := range projects {
		data, err := svc.Services.List("projects/" + project.ID).Do()
		if err != nil {
			log.Println(err)
			return listOfAPIs, err
		}

		for _, service := range data.Services {
			enabled := true
			if service.State == "DISABLED" {
				enabled = false
			}
			listOfAPIs = append(listOfAPIs, API{
				Namespace: service.Config.Name,
				Title:     service.Config.Title,
				Enabled:   enabled,
			})
		}
	}
	return listOfAPIs, nil
}
