package civo

import (
	"github.com/civo/civogo"
)

type Civo struct{}

var client *civogo.Client

func getCivoClient(apiKey, regionCode string) (*civogo.Client, error) {
	if client != nil {
		return client, nil
	}
	if apiKey == "" {
		panic("apiKey provided for connecting to Civo client is empty")
	} else if regionCode == "" {
		panic("regionCode provided for connecting to Civo client is empty")
	}
	client, err := civogo.NewClient(apiKey, regionCode)
	return client, err
}
