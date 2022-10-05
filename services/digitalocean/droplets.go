package digitalocean

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
)

type Droplet struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Tags   []string `json:"tags"`
	Image  string   `json:"image"`
	Region string   `json:"region"`
	Status string   `json:"status"`
	Disk   int      `json:"disk"`
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
			Name:   droplet.Name,
			ID:     fmt.Sprintf("%d", droplet.ID),
			Tags:   droplet.Image.Tags,
		})
	}
	return listOfDroplets, nil
}
