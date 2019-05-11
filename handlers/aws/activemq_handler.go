package aws

import (
	"net/http"
)

func (handler *AWSHandler) ActiveMQBrokersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_mq_brokers")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListBrokers(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "mq:ListBrokers is missing")
		} else {
			handler.cache.Set("aws_mq_brokers", response)
			respondWithJSON(w, 200, response)
		}
	}
}
