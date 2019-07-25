package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

func (dg DigitalOcean) DescribeProjects(client *godo.Client) (int, error) {
	projects, _, err := client.Projects.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return 0, err
	}
	return len(projects), nil
}
