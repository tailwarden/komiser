package handlers

import (
	"fmt"
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) LambdaFunctionHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("lambda_functions")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeLambdaFunctions(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "lambda:ListFunctions is missing")
		} else {
			handler.cache.Set("lambda_functions", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) GetLambdaInvocationMetrics(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("lambda_invocations")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetLambdaInvocationMetrics(handler.cfg)
		fmt.Println(err)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "lambda:ListFunctions is missing")
		} else {
			handler.cache.Set("lambda_invocations", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
