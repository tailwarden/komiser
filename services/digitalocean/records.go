package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

type Record struct {
	A     int `json:"a"`
	CNAME int `json:"cname"`
}

func (dg DigitalOcean) DescribeRecords(client *godo.Client) (Record, error) {
	record := Record{}

	domains, _, err := client.Domains.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return record, err
	}

	for _, domain := range domains {
		records, _, err := client.Domains.Records(context.TODO(), domain.Name, &godo.ListOptions{})
		if err != nil {
			return record, err
		}

		for _, r := range records {
			if r.Type == "A" {
				record.A++
			}
			if r.Type == "CNAME" {
				record.A++
			}
		}
	}
	return record, nil
}
