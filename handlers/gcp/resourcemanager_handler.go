package aws

import (
	"net/http"
)

func (handler *GCPHandler) ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_resourcemanager_projects")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetProjects()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudresourcemanager:CloudPlatformReadOnlyScope is missing")
		} else {
			handler.cache.Set("gcp_resourcemanager_projects", response)
			respondWithJSON(w, 200, response)
		}
	}
}
