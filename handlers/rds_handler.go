package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) RDSInstancePerEngineHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("rds_engine")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeRDSInstancesPerEngine(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("rds_engine", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
