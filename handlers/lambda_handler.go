package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) LambdaPerRuntimeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("lambda_runtime")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeLambdaFunctionsPerRuntime(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("lambda_runtime", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
