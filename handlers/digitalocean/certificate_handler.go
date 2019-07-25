package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) CertificatesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.certificates")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeCertificates(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.certificates", response)
			respondWithJSON(w, 200, response)
		}
	}
}
