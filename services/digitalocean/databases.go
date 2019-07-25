package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

func (dg DigitalOcean) DescribeDatabases(client *godo.Client) (int, error) {
	databases, _, err := client.Databases.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return 0, err
	}
	return len(databases), nil
}
