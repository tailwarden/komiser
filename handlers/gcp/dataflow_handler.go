package aws

import (
	"net/http"
)

func (handler *GCPHandler) DataflowJobsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_dataflow_jobs")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetDataflowJobs()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "dataflow:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("gcp_dataflow_jobs", response)
			respondWithJSON(w, 200, response)
		}
	}
}
