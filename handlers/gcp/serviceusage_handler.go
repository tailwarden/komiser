package aws

import (
	"net/http"
)

func (handler *GCPHandler) EnabledAPIsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("enabled_apis")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetEnabledAPIs()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "serviceusage:CloudPlatformReadOnlyScope is missing")
		} else {
			handler.cache.Set("enabled_apis", response)
			respondWithJSON(w, 200, response)
		}
	}
}
