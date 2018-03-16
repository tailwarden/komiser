package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) IAMRolesTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("role_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeIAMRolesTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("role_total", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) IAMGroupsTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("group_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeIAMGroupsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("group_total", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) IAMPoliciesTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("policy_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeIAMPoliciesTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("policy_total", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) IAMUsersTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("user_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeIAMUsersTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("user_total", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
