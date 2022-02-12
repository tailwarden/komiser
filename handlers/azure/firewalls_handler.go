package azure

import (
	"fmt"
	"net/http"

	"github.com/mlabouardy/komiser/handlers/azure/config"
)

func (handler *AzureHandler) FirewallsHandler(w http.ResponseWriter, r *http.Request) {
	err := config.ParseEnvironment()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse environment")
	}
	subscriptionID := config.SubscriptionID()
	key := fmt.Sprintf("azure.%s.security.firewalls", subscriptionID)
	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.azure.GetFirewallsCount(subscriptionID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Security:GetFirewallsCount is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, 200, response)
		}
	}
}
