package aws

import (
	"net/http"
)

func (handler *AWSHandler) CloudTrailConsoleSignInEventsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_events_sign_in")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.CloudTrailConsoleSignInEvents(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudtrail:LookupEvents is missing")
		} else {
			handler.cache.Set("aws_events_sign_in", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) CloudTrailConsoleSignInSourceIpEventsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_events_source_ip")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.CloudTrailConsoleSignInSourceIpEvents(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudtrail:LookupEvents is missing")
		} else {
			handler.cache.Set("aws_events_source_ip", response)
			respondWithJSON(w, 200, response)
		}
	}
}
