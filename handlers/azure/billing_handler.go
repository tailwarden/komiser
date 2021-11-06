package azure

import (
	"fmt"
	"net/http"

	"github.com/mlabouardy/komiser/handlers/azure/config"
)

func (handler *AzureHandler) CostManagementHandler(w http.ResponseWriter, r *http.Request) {
	err := config.ParseEnvironment()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse environment")
	}
	subscriptionID := config.SubscriptionID()
	key := fmt.Sprintf("azure.%s.billing.bill", subscriptionID)
	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.azure.Usage(subscriptionID)
		if err != nil {
			fmt.Println("error:", err)
			respondWithError(w, http.StatusInternalServerError, "Bill:Usage is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, 200, response)
		}
	}
}
