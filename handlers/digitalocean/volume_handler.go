package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) VolumesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.volumes")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeVolumes(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.volumes", response)
			respondWithJSON(w, 200, response)
		}
	}
}
