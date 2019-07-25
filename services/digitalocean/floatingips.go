package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

func (dg DigitalOcean) DescribeFloatingIps(client *godo.Client) (int, error) {
	ips, _, err := client.FloatingIPs.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return 0, err
	}
	return len(ips), nil
}
