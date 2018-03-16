package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) CostAndUsageHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cost_usage")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCostAndUsage(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("cost_usage", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
