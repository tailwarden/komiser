package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

func (dg DigitalOcean) DescribeCDN(client *godo.Client) (int, error) {
	cdns, _, err := client.CDNs.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return 0, err
	}
	return len(cdns), nil
}
