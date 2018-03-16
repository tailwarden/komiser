package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) CloudWatchAlarmsStateHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cloudwatch_alarm")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCloudWatchAlarmsPerState(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("cloudwatch_alarm", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
