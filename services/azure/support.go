package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/support/mgmt/2020-04-01/support"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getTicketsClient(subscriptionID string) support.TicketsClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	ticketsClient := support.NewTicketsClient(subscriptionID)
	ticketsClient.Authorizer = a
	return ticketsClient
}

func (azure Azure) DescribeTickets(subscriptionID string) (map[string]interface{}, error) {
	closedTickets := make(map[string]int, 0)
	resolvedTickets := make(map[string]int, 0)
	ticketsClient := getTicketsClient(subscriptionID)
	ctx := context.Background()
	var pageSize *int32
	for resultItr, err := ticketsClient.ListComplete(ctx, pageSize, ""); resultItr.NotDone(); resultItr.Next() {
		if err != nil {
			return map[string]interface{}{}, err
		}
		ticket := resultItr.Value()
		
	}
	


}
