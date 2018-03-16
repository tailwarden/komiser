package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) HostedZoneTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("hosted_zone_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeHostedZonesTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("hosted_zone_total", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
