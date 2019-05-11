package aws

import (
	"net/http"
)

func (handler *GCPHandler) LoggingBillableReceivedBytesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_logging_bytes")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetBillableReceivedLogs()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "iam:CloudPlatformScope is missing")
		} else {
			handler.cache.Set("gcp_logging_bytes", response)
			respondWithJSON(w, 200, response)
		}
	}
}
