package aws

import (
	"net/http"
)

func (handler *GCPHandler) CloudFunctionsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_cloud_functions")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.CloudFunctions()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudfunctions:CloudPlatformScope is missing")
		} else {
			handler.cache.Set("gcp_cloud_functions", response)
			respondWithJSON(w, 200, response)
		}
	}
}
