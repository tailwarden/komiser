package aws

import (
	"net/http"
)

func (handler *AWSHandler) SQSQueuesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_sqs")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeQueues(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "sqs:ListQueues is missing")
		} else {
			handler.cache.Set("aws_sqs", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) GetNumberOfMessagesSentAndDeletedSQSHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_sqs_messages")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetNumberOfMessagesSentAndDeletedSQS(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set("aws_sqs_messages", response)
			respondWithJSON(w, 200, response)
		}
	}
}
