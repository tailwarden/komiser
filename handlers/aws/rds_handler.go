package aws

import (
	"net/http"
)

func (handler *AWSHandler) RDSInstanceHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_rds")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeRDSInstances(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "rds:DescribeDBInstances is missing")
		} else {
			handler.cache.Set("aws_rds", response)
			respondWithJSON(w, 200, response)
		}
	}
}
