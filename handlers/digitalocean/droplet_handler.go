package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) DropletsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.droplets")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeDroplets(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.droplets", response)
			respondWithJSON(w, 200, response)
		}
	}
}
