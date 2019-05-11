package aws

import (
	"net/http"
)

func (handler *AWSHandler) APIGatewayListCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_acm_certificates")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListCertificates(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "acm:ListCertificates is missing")
		} else {
			handler.cache.Set("aws_acm_certificates", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) APIGatewayExpiredCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_acm_expired")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListExpiredCertificates(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "acm:ListCertificates is missing")
		} else {
			handler.cache.Set("aws_acm_expired", response)
			respondWithJSON(w, 200, response)
		}
	}
}
