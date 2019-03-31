package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) UsedRegionsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("used_regions")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeResources(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "resourcegroupstaggingapi:get-resources is missing")
		} else {
			handler.cache.Set("used_regions", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
