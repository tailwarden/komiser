package aws

import (
	"net/http"
)

func (handler *AWSHandler) ESListDomainsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_es_domains")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListESDomains(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "es:ListDomainNames is missing")
		} else {
			handler.cache.Set("aws_es_domains", response)
			respondWithJSON(w, 200, response)
		}
	}
}
