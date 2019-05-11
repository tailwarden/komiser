package aws

import (
	"net/http"
)

func (handler *AWSHandler) S3BucketsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_s3")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeS3Buckets(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "s3:ListAllMyBuckets is missing")
		} else {
			handler.cache.Set("aws_s3", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) S3BucketsSizeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_s3_size")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetBucketsSize(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set("aws_s3_size", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) S3BucketsObjectsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_s3_objects")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetBucketsObjects(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set("aws_s3_objects", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) GetEmptyBucketsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_s3_empty")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetEmptyBuckets(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set("aws_s3_empty", response)
			respondWithJSON(w, 200, response)
		}
	}
}
