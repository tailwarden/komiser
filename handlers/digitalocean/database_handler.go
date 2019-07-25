package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) DatabasesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.databases")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeDatabases(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.databases", response)
			respondWithJSON(w, 200, response)
		}
	}
}
