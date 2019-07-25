package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

func (dg DigitalOcean) DescribeLoadBalancers(client *godo.Client) (int, error) {
	loadbalancers, _, err := client.LoadBalancers.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return 0, err
	}
	return len(loadbalancers), nil
}
