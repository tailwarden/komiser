package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) RDSInstanceHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("rds")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeRDSInstances(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "rds:DescribeDBInstances is missing")
		} else {
			handler.cache.Set("rds", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
