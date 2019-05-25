package aws

import (
	"net/http"
)

func (handler *GCPHandler) DataprocJobsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_dataproc_jobs")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetDataprocJobs()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "dataproc:CloudPlatformScope is missing")
		} else {
			handler.cache.Set("gcp_dataproc_jobs", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) DataprocClustersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_dataproc_clusters")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetDataprocClusters()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "dataproc:CloudPlatformScope is missing")
		} else {
			handler.cache.Set("gcp_dataproc_clusters", response)
			respondWithJSON(w, 200, response)
		}
	}
}
