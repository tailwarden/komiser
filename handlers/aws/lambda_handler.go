package aws

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func (handler *AWSHandler) LambdaFunctionHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg, err := external.LoadDefaultAWSConfig()

	if handler.multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	key := fmt.Sprintf("aws.%s.lambda.functions", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeLambdaFunctions(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "lambda:ListFunctions is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) GetLambdaInvocationMetrics(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg, err := external.LoadDefaultAWSConfig()

	if handler.multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	key := fmt.Sprintf("aws.%s.lambda.invocations", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetLambdaInvocationMetrics(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) GetLambdaErrorsMetrics(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg, err := external.LoadDefaultAWSConfig()

	if handler.multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	key := fmt.Sprintf("aws.%s.lambda.errors", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetLambdaErrorsMetrics(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, 200, response)
		}
	}
}
