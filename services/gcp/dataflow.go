package gcp

import (
	"context"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	dataflow "google.golang.org/api/dataflow/v1b3"
)

func (gcp GCP) GetDataflowJobs() (int, error) {
	sum := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, dataflow.ComputeReadonlyScope)
	if err != nil {
		return sum, err
	}

	client := oauth2.NewClient(context.Background(), src)

	svc, err := dataflow.New(client)
	if err != nil {
		return sum, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return sum, err
	}

	for _, project := range projects {
		jobs, err := svc.Projects.Jobs.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return sum, err
		}
		sum += len(jobs.Jobs)
	}

	return sum, nil
}
