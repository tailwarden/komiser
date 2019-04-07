package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) EC2InstancesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ec2")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeInstances(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeInstances is missing")
		} else {
			handler.cache.Set("ec2", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) AutoScalingGroupHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("asg")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeAutoScalingGroups(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "autoscaling:DescribeAutoScalingGroups is missing")
		} else {
			handler.cache.Set("asg", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) ListUnrestrictedSecurityGroups(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("sg_unrestricted")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListUnrestrictedSecurityGroups(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("sg_unrestricted", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
