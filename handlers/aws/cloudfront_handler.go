package aws

import (
	"net/http"
)

func (handler *AWSHandler) CloudFrontDistributionsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_cloudfront")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCloudFrontDistributions(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudfront:ListDistributions is missing")
		} else {
			handler.cache.Set("aws_cloudfront", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) CloudFrontRequestsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_cloudfront_requests")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetCloudFrontRequests(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set("aws_cloudfront_requests", response)
			respondWithJSON(w, 200, response)
		}
	}
}
