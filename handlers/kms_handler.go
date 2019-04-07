package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) KMSKeysHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("kms_keys")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListKeys(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "kms:list-keys is missing")
		} else {
			handler.cache.Set("kms_keys", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
