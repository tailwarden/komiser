package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) S3TotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("s3_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeS3BucketsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("s3_total", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
