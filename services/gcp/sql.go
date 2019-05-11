package gcp

import (
	"context"
	"log"

	. "github.com/mlabouardy/komiser/models/gcp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

func (gcp GCP) GetSqlInstances() ([]SqlInstance, error) {
	listOfInstances := make([]SqlInstance, 0)

	src, err := google.DefaultTokenSource(oauth2.NoContext, sqladmin.CloudPlatformScope)
	if err != nil {
		return listOfInstances, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := sqladmin.New(client)
	if err != nil {
		return listOfInstances, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return listOfInstances, err
	}

	for _, project := range projects {
		instances, err := svc.Instances.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return listOfInstances, err
		}

		for _, instance := range instances.Items {
			listOfInstances = append(listOfInstances, SqlInstance{
				Region: instance.Region,
				State:  instance.State,
				Kind:   instance.DatabaseVersion,
			})
		}
	}

	return listOfInstances, nil
}
