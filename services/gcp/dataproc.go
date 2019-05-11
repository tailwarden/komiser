package gcp

import (
	"context"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	dataproc "google.golang.org/api/dataproc/v1"
)

func (gcp GCP) GetDataprocJobs() (int, error) {
	sum := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, dataproc.CloudPlatformScope)
	if err != nil {
		return sum, err
	}

	client := oauth2.NewClient(context.Background(), src)

	svc, err := dataproc.New(client)
	if err != nil {
		return sum, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return sum, err
	}

	regions, err := gcp.GetRegions()
	if err != nil {
		return sum, err
	}

	for _, project := range projects {
		for _, region := range regions {
			jobs, err := svc.Projects.Regions.Jobs.List(project.ID, region).Do()
			if err != nil {
				log.Println(err)
				return sum, err
			}
			sum += len(jobs.Jobs)
		}
	}

	return sum, nil
}

func (gcp GCP) GetDataprocClusters() (int, error) {
	sum := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, dataproc.CloudPlatformScope)
	if err != nil {
		return sum, err
	}

	client := oauth2.NewClient(context.Background(), src)

	svc, err := dataproc.New(client)
	if err != nil {
		return sum, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return sum, err
	}

	regions, err := gcp.GetRegions()
	if err != nil {
		return sum, err
	}

	for _, project := range projects {
		for _, region := range regions {
			data, err := svc.Projects.Regions.Clusters.List(project.ID, region).Do()
			if err != nil {
				log.Println(err)
				return sum, err
			}
			sum += len(data.Clusters)
		}
	}

	return sum, nil
}
