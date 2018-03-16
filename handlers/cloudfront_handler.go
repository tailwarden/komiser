package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) CloudFrontDistributionsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cloudfront")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCloudFrontDistributions(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudfront:ListDistributions is missing")
		} else {
			handler.cache.Set("cloudfront", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
