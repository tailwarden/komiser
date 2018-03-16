package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) CloudFrontDistributionsTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cloudfront_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCloudFrontDistributionsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("cloudfront_total", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
