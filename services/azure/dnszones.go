package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2018-05-01/dns"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getDNSZonesClient(subscriptionID string) dns.ZonesClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	dnsZonesClient := dns.NewZonesClient(subscriptionID)
	dnsZonesClient.Authorizer = a
	return dnsZonesClient
}

func (azure Azure) GetDNSZonesCount(subscriptionID string) (int, error) {
	var dnsZonesCount int
	var top *int32
	dnsZonesClient := getDNSZonesClient(subscriptionID)
	ctx := context.Background()

	for dnsItr, err := dnsZonesClient.ListComplete(ctx, top); dnsItr.NotDone(); dnsItr.Next() {
		if err != nil {
			return dnsZonesCount, err
		}
		dnsZonesCount++
	}
	return dnsZonesCount, nil
}
