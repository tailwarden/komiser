package aws

import (
	"net/http"
)

func (handler *AWSHandler) KinesisListStreamsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_kinesis_streams")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListStreams(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "kinesis:ListStreams is missing")
		} else {
			handler.cache.Set("aws_kinesis_streams", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) KinesisListShardsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_kinesis_shards")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListShards(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "kinesis:ListShards is missing")
		} else {
			handler.cache.Set("aws_kinesis_shards", response)
			respondWithJSON(w, 200, response)
		}
	}
}
