package aws

import (
	"net/http"
)

func (handler *GCPHandler) AppEngineOutgoingBandwidthHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_app_engine_bandwidth")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetAppEngineOutgoingBandwidth()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "monitoring:MonitoringReadScope is missing")
		} else {
			handler.cache.Set("gcp_app_engine_bandwidth", response)
			respondWithJSON(w, 200, response)
		}
	}
}
