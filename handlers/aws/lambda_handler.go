package aws

import (
	"net/http"
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
			handler.cache.Set("lambda_functions", response)
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
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set("lambda_invocations", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) GetLambdaErrorsMetrics(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("lambda_errors")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetLambdaErrorsMetrics(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set("lambda_errors", response)
			respondWithJSON(w, 200, response)
		}
	}
}
