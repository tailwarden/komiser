package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) EBSSizeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ebs_size")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeVolumesTotalSize(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("ebs_size", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) EBSFamilyHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ebs_family")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeVolumesPerFamily(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("ebs_family", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) EBSStateHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ebs_state")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeVolumesPerState(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("ebs_state", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
