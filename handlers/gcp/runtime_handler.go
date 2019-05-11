package aws

import (
	"net/http"
)

func (handler *GCPHandler) ConsumedAPIRequestsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_api_consumed_requests")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetConsumedAPIRequests()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "monitoring:MonitoringReadScope is missing")
		} else {
			handler.cache.Set("gcp_api_consumed_requests", response)
			respondWithJSON(w, 200, response)
		}
	}
}
