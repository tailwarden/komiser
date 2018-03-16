package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) EBSHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ebs")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeVolumes(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeVolumes is missing")
		} else {
			handler.cache.Set("ebs", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
