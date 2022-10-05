package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

type Snapshot struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Size   float64  `json:"size"`
	Region string   `json:"region"`
	Tags   []string `json:"tags"`
}

func (dg DigitalOcean) DescribeSnapshots(client *godo.Client) ([]Snapshot, error) {
	listOfSnapshots := make([]Snapshot, 0)

	snapshots, _, err := client.Snapshots.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return listOfSnapshots, err
	}

	for _, snapshot := range snapshots {
		listOfSnapshots = append(listOfSnapshots, Snapshot{
			Size:   snapshot.SizeGigaBytes,
			ID:     snapshot.ID,
			Name:   snapshot.Name,
			Region: snapshot.Regions[0],
			Tags:   snapshot.Tags,
		})
	}

	return listOfSnapshots, nil
}
