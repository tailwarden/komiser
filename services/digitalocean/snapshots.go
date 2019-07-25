package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

type Snapshot struct {
	Size float64 `json:"size"`
}

func (dg DigitalOcean) DescribeSnapshots(client *godo.Client) ([]Snapshot, error) {
	listOfSnapshots := make([]Snapshot, 0)

	snapshots, _, err := client.Snapshots.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return listOfSnapshots, err
	}

	for _, snapshot := range snapshots {
		listOfSnapshots = append(listOfSnapshots, Snapshot{
			Size: snapshot.SizeGigaBytes,
		})
	}

	return listOfSnapshots, nil
}
