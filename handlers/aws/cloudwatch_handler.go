package aws

import (
	"net/http"
)

func (handler *AWSHandler) CloudWatchAlarmsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cloudwatch_alarms")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCloudWatchAlarms(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:DescribeAlarms is missing")
		} else {
			handler.cache.Set("cloudwatch_alarms", response)
			respondWithJSON(w, 200, response)
		}
	}
}
