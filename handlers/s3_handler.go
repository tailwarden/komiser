package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) S3BucketsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("s3")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeS3Buckets(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "s3:ListAllMyBuckets is missing")
		} else {
			handler.cache.Set("s3", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
