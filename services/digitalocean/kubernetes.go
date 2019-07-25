package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

type Kubernetes struct {
	Nodes  int    `json:"nodes"`
	Region string `json:"region"`
}

func (dg DigitalOcean) DescribeK8s(client *godo.Client) ([]Kubernetes, error) {
	listOfClusters := make([]Kubernetes, 0)

	clusters, _, err := client.Kubernetes.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return listOfClusters, err
	}

	for _, cluster := range clusters {
		nodes := 0
		for _, pool := range cluster.NodePools {
			nodes += pool.Count
		}
		listOfClusters = append(listOfClusters, Kubernetes{
			Region: cluster.RegionSlug,
			Nodes:  nodes,
		})
	}
	return listOfClusters, nil
}
