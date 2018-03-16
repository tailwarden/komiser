package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) CloudWatchAlarmsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cloudwatch")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCloudWatchAlarms(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:DescribeAlarms is missing")
		} else {
			handler.cache.Set("cloudwatch", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
