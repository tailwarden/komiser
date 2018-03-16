package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) LambdaFunctionHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("lambda")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeLambdaFunctions(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "lambda:ListFunctions is missing")
		} else {
			handler.cache.Set("lambda", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
