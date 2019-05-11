package aws

import (
	"net/http"
)

func (handler *GCPHandler) BillingLastSixMonthsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_billing_history")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.CostInLastSixMonths()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "bigquery:Query is missing")
		} else {
			handler.cache.Set("gcp_billing_history", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) BillingPerServiceHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_billing_per_service")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.MonthlyCostPerService()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "bigquery:Query is missing")
		} else {
			handler.cache.Set("gcp_billing_per_service", response)
			respondWithJSON(w, 200, response)
		}
	}
}
