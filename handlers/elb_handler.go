package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) ElasticLoadBalancerFamilyHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("elb_family")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeElasticLoadBalancerPerFamily(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("elb_family", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
