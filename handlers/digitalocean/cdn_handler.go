package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) ContentDeliveryNetworksHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.cdn")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeCDN(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.cdn", response)
			respondWithJSON(w, 200, response)
		}
	}
}
