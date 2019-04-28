package aws

import (
	"net/http"
)

func (handler *AWSHandler) SNSTopicsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("sns")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeSNSTopics(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "sns:ListTopics is missing")
		} else {
			handler.cache.Set("sns", response)
			respondWithJSON(w, 200, response)
		}
	}
}
