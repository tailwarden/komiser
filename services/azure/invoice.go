package azure

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/preview/billing/mgmt/2020-05-01-preview/billing"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	. "github.com/mlabouardy/komiser/models/azure"
)

func getInvoicesClient(subscriptionID string) (billing.InvoicesClient, error) {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return billing.InvoicesClient{}, err
	}
	invoicesClient := billing.NewInvoicesClient(subscriptionID)
	invoicesClient.Authorizer = a
	return invoicesClient, nil
}

func (azure Azure) GetBilling(subscriptionID string) (Invoice, error) {
	var bill float64
	var currency string
	invoicesClient, err := getInvoicesClient(subscriptionID)
	if err != nil {
		return Invoice{}, err
	}
	ctx := context.Background()
	current, previous := getTargetDates()
	for invItr, err := invoicesClient.ListByBillingSubscriptionComplete(ctx, previous, current); invItr.NotDone(); invItr.Next() {
		if err != nil {
			return Invoice{}, err
		}
		bill += *invItr.Value().InvoiceProperties.TotalAmount.Value
		currency = *invItr.Value().InvoiceProperties.TotalAmount.Currency
	}
	return Invoice{
		Amount:   bill,
		Currency: currency,
	}, nil
}

// Expected dates are in the format MM-DD-YYYY
func getTargetDates() (string, string) {
	currentTime := time.Now()
	monthMap := getMonthMap()
	currentDate := monthMap[string(currentTime.Month())] + "-" + string(currentTime.Day()) + "-" + string(currentTime.Year())
	oldTime := currentTime.AddDate(0, -3, 0)
	oldDate := monthMap[string(oldTime.Month())] + "-" + string(oldTime.Day()) + "-" + string(oldTime.Year())
	return currentDate, oldDate
}

func getMonthMap() map[string]string {
	months := make(map[string]string)
	months["January"] = "01"
	months["February"] = "02"
	months["March"] = "03"
	months["April"] = "04"
	months["May"] = "05"
	months["June"] = "06"
	months["July"] = "07"
	months["August"] = "08"
	months["September"] = "09"
	months["October"] = "10"
	months["November"] = "11"
	months["December"] = "12"
	return months
}
