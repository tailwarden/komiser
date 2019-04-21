package aws

import (
	"net/http"
)

func (handler *AWSHandler) SnapshotHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("snapshot")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeSnapshots(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeSnapshots is missing")
		} else {
			handler.cache.Set("snapshot", response)
			respondWithJSON(w, 200, response)
		}
	}
}
