package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

func (dg DigitalOcean) DescribeSSHKeys(client *godo.Client) (int, error) {
	keys, _, err := client.Keys.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return 0, err
	}
	return len(keys), nil
}
