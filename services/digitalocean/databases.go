package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

type DigitalOceanDatabase struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Region string   `json:"region"`
	Tags   []string `json:"tags"`
}

func (dg DigitalOcean) DescribeDatabases(client *godo.Client) ([]DigitalOceanDatabase, error) {
	listOfDbs := make([]DigitalOceanDatabase, 0)
	databases, _, err := client.Databases.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return listOfDbs, err
	}

	for _, db := range databases {
		listOfDbs = append(listOfDbs, DigitalOceanDatabase{
			ID:     db.ID,
			Name:   db.Name,
			Region: db.RegionSlug,
			Tags:   db.Tags,
		})
	}
	return listOfDbs, nil
}
