package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) FloatingIpsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.floatingips")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeFloatingIps(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.floatingips", response)
			respondWithJSON(w, 200, response)
		}
	}
}
