package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) SnapshotsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.snapshots")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeSnapshots(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.snapshots", response)
			respondWithJSON(w, 200, response)
		}
	}
}
