package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) DynamoDBTableTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("dynamodb_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeDynamoDBTablesTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("dynamodb_total", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) DynamoDBProvisionedThroughputHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("dynamodb_throughput")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeDynamoDBTablesProvisionedThroughput(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("dynamodb_throughput", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
