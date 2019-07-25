package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) RecordsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.records")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeRecords(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.records", response)
			respondWithJSON(w, 200, response)
		}
	}
}
