package aws

import (
	"net/http"
)

func (handler *AWSHandler) APIGatewayRequestsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("apigateway_requests")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetAPIGatewayRequests(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set("apigateway_requests", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) APIGatewayRestAPIsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("apigateway_apis")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetRestAPIs(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "apigateway:GET is missing")
		} else {
			handler.cache.Set("apigateway_apis", response)
			respondWithJSON(w, 200, response)
		}
	}
}
