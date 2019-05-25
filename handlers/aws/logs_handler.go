package aws

import (
	"net/http"
)

func (handler *AWSHandler) LogsVolumeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_logs_volume")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetLogsVolume(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set("aws_logs_volume", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) MaximumLogsRetentionPeriodHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_logs_retention")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.MaximumLogsRetentionPeriod(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "logs:DescribeLogGroups is missing")
		} else {
			handler.cache.Set("aws_logs_retention", response)
			respondWithJSON(w, 200, response)
		}
	}
}
