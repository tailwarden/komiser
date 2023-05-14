package gcpcomputepricing

import (
	"encoding/json"
	"net/http"
)

var (
	url = "https://www.gstatic.com/cloud-site-ux/pricing/data/gcp-compute.json"

	httpClient = http.DefaultClient
)

func Fetch() (*Pricing, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var pricing Pricing
	if err := json.NewDecoder(resp.Body).Decode(&pricing); err != nil {
		return nil, err
	}
	return &pricing, nil
}

func SetHTTPClient(client *http.Client) {
	httpClient = client
}

func SetURL(u string) {
	url = u
}
