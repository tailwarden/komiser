package aws

import (
	"net/http"
)

func (handler *AWSHandler) LogsVolumeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("logs_volume")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetLogsVolume(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set("logs_volume", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) MaximumLogsRetentionPeriodHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("logs_retention")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.MaximumLogsRetentionPeriod(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatchlogs:DescribeLogGroups is missing")
		} else {
			handler.cache.Set("logs_retention", response)
			respondWithJSON(w, 200, response)
		}
	}
}
