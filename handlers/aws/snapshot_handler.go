package aws

import (
	"net/http"
)

func (handler *AWSHandler) SnapshotHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_snapshot")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeSnapshots(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeSnapshots is missing")
		} else {
			handler.cache.Set("aws_snapshot", response)
			respondWithJSON(w, 200, response)
		}
	}
}
