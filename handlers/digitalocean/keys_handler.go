package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) SSHKeysHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.keys")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeSSHKeys(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.keys", response)
			respondWithJSON(w, 200, response)
		}
	}
}
