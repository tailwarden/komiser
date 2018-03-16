package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) SnapshotTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("snapshot_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeSnapshotsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("snapshot_total", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) SnapshotSizeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("snapshot_size")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeSnapshotsSize(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("snapshot_size", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
