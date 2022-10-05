package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

type Volume struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Region string   `json:"region"`
	Tags   []string `json:"tags"`
	Size   int64    `json:"size"`
}

func (dg DigitalOcean) DescribeVolumes(client *godo.Client) ([]Volume, error) {
	listOfVolumes := make([]Volume, 0)

	volumes, _, err := client.Storage.ListVolumes(context.TODO(), &godo.ListVolumeParams{})
	if err != nil {
		return listOfVolumes, err
	}

	for _, v := range volumes {
		listOfVolumes = append(listOfVolumes, Volume{
			Size:   v.SizeGigaBytes,
			ID:     v.ID,
			Name:   v.Name,
			Region: v.Region.Name,
			Tags:   v.Tags,
		})
	}

	return listOfVolumes, nil
}
