package azure

import (
	"fmt"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/config"
	"net/http"
)

func (handler *AzureHandler) APIGatewayListCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	err := config.ParseEnvironment()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse environment")
	}
	subscriptionID := config.SubscriptionID()
	key := fmt.Printf("azure.%s.acm.certificates", subscriptionID)
	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.azure.ListCertificates(subscriptionID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "acm:ListCertificates is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AzureHandler) APIGatewayListExpiredCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	err := config.ParseEnvironment()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse environment")
	}
	subscriptionID := config.SubscriptionID()
	key := fmt.Printf("azure.%s.acm.expired", subscriptionID)
	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.azure.ListExpiredCertificates(subscriptionID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "acm:ListCertificates is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, 200, response)
		}
	}
}
