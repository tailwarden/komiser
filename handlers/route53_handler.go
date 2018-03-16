package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) HostedZoneHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("hosted_zone")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeHostedZones(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "route53:ListHostedZones is missing")
		} else {
			handler.cache.Set("hosted_zone", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
