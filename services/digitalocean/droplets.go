package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

type Droplet struct {
	Image  string `json:"image"`
	Region string `json:"region"`
	Status string `json:"status"`
	Disk   int    `json:"disk"`
}

func (dg DigitalOcean) DescribeDroplets(client *godo.Client) ([]Droplet, error) {
	listOfDroplets := make([]Droplet, 0)

	droplets, _, err := client.Droplets.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return listOfDroplets, err
	}

	for _, droplet := range droplets {
		listOfDroplets = append(listOfDroplets, Droplet{
			Disk:   droplet.Disk,
			Image:  droplet.Image.Distribution,
			Status: droplet.Status,
			Region: droplet.Region.Slug,
		})
	}
	return listOfDroplets, nil
}
