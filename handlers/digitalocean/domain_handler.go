package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) DomainsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.domains")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeDomains(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.domains", response)
			respondWithJSON(w, 200, response)
		}
	}
}
