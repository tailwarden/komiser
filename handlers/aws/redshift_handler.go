package aws

import (
	"net/http"
)

func (handler *AWSHandler) DescribeRedshiftClustersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_redshift_clusters")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeRedshiftClusters(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "redshift:DescribeClusters is missing")
		} else {
			handler.cache.Set("aws_redshift_clusters", response)
			respondWithJSON(w, 200, response)
		}
	}
}
