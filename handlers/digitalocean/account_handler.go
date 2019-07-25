package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) AccountProfileHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.account")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeAccount(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.account", response)
			respondWithJSON(w, 200, response)
		}
	}
}
