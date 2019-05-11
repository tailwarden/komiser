package aws

import (
	"net/http"
)

func (handler *AWSHandler) SWFListDomainsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_swf_domains")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetSWFDomains(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "swf:ListDomains is missing")
		} else {
			handler.cache.Set("aws_swf_domains", response)
			respondWithJSON(w, 200, response)
		}
	}
}
