package aws

import (
	"net/http"
)

func (handler *AWSHandler) EKSClustersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("eks_clusters")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeEKSClusters(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "eks:ListClusters is missing")
		} else {
			handler.cache.Set("eks_clusters", response)
			respondWithJSON(w, 200, response)
		}
	}
}
