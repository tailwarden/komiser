package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) EC2RegionHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ec2_region")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeInstancesPerRegion(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("ec2_region", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) EC2FamilyHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ec2_family")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeInstancesPerFamily(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("ec2_family", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) EC2StateHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ec2_state")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeInstancesPerState(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("ec2_state", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) AutoScalingGroupTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("asg_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeAutoScalingGroupsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("asg_total", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
