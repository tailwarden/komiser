package aws

import (
	"fmt"
	"net/http"
)

func (handler *GCPHandler) BillingLastSixMonthsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("billing_history")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.CostInLastSixMonths()
		if err != nil {
			fmt.Println(err)
			respondWithError(w, http.StatusInternalServerError, "bigquery:Query is missing")
		} else {
			handler.cache.Set("billing_history", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) BillingPerServiceHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("billing_per_service")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.MonthlyCostPerService()
		if err != nil {
			fmt.Println(err)
			respondWithError(w, http.StatusInternalServerError, "bigquery:Query is missing")
		} else {
			handler.cache.Set("billing_per_service", response)
			respondWithJSON(w, 200, response)
		}
	}
}
