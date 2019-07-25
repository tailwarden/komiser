package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.projects")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeProjects(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.projects", response)
			respondWithJSON(w, 200, response)
		}
	}
}
