package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) KubernetesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.k8s")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeK8s(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.k8s", response)
			respondWithJSON(w, 200, response)
		}
	}
}
