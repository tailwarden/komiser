package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

func (dg DigitalOcean) DescribeDomains(client *godo.Client) (int, error) {
	domains, _, err := client.Domains.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return 0, err
	}
	return len(domains), nil
}
