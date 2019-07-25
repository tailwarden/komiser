package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) ActionsHistoryHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.actions")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeActions(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.actions", response)
			respondWithJSON(w, 200, response)
		}
	}
}
