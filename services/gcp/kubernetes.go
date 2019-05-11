package gcp

import (
	"context"
	"fmt"
	"log"

	. "github.com/mlabouardy/komiser/models/gcp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	container "google.golang.org/api/container/v1"
)

func (gcp GCP) GetKubernetesClusters() ([]Kubernetes, error) {
	listOfClusters := make([]Kubernetes, 0)

	src, err := google.DefaultTokenSource(oauth2.NoContext, container.CloudPlatformScope)
	if err != nil {
		return listOfClusters, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := container.New(client)
	if err != nil {
		return listOfClusters, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return listOfClusters, err
	}

	for _, project := range projects {
		uri := fmt.Sprintf("projects/%s/locations/-", project.ID)
		clusters, err := svc.Projects.Locations.Clusters.List(uri).Do()
		if err != nil {
			log.Println(err)
			return listOfClusters, err
		}

		for _, cluster := range clusters.Clusters {
			listOfClusters = append(listOfClusters, Kubernetes{
				MachineType: cluster.NodeConfig.MachineType,
				Status:      cluster.Status,
				Zone:        cluster.Zone,
				Nodes:       cluster.CurrentNodeCount,
			})
		}
	}
	return listOfClusters, nil
}
