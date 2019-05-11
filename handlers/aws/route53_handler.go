package aws

import (
	"net/http"
)

func (handler *AWSHandler) Route53HostedZonesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_route53_zones")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeHostedZones(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "route53:ListHostedZones is missing")
		} else {
			handler.cache.Set("aws_route53_zones", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) Route53ARecordsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_route53_a_records")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeARecords(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "route53:ListResourceRecordSets is missing")
		} else {
			handler.cache.Set("aws_route53_a_records", response)
			respondWithJSON(w, 200, response)
		}
	}
}
