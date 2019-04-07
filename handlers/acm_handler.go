package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) APIGatewayListCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("acm_certificates")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListCertificates(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "acm:ListDistributions is missing")
		} else {
			handler.cache.Set("acm_certificates", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) APIGatewayExpiredCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("acm_expired")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListExpiredCertificates(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "acm:ListDistributions is missing")
		} else {
			handler.cache.Set("acm_expired", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
