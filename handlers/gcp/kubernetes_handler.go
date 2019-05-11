package aws

import (
	"net/http"
)

func (handler *GCPHandler) KubernetesClustersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_kubernetes_clusters")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetKubernetesClusters()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "container:CloudPlatformScope is missing")
		} else {
			handler.cache.Set("gcp_kubernetes_clusters", response)
			respondWithJSON(w, 200, response)
		}
	}
}
