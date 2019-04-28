package aws

import (
	"net/http"
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
			handler.cache.Set("ec2", response)
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
			handler.cache.Set("asg", response)
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
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeSecurityGroups is missing")
		} else {
			handler.cache.Set("sg_unrestricted", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) ScheduledEC2Instances(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ec2_scheduled")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeScheduledInstances(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeScheduledInstances is missing")
		} else {
			handler.cache.Set("ec2_scheduled", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) SpotEC2Instances(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ec2_spot")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeSpotInstances(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeSpotFleetRequests is missing")
		} else {
			handler.cache.Set("ec2_spot", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) ReservedEC2Instances(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ec2_reserved")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeReservedInstances(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeReservedInstances is missing")
		} else {
			handler.cache.Set("ec2_reserved", response)
			respondWithJSON(w, 200, response)
		}
	}
}
