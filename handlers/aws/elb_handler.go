package aws

import (
	"net/http"
)

func (handler *AWSHandler) ElasticLoadBalancerHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("elb")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeElasticLoadBalancer(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "elasticloadbalancing:DescribeLoadBalancers is missing")
		} else {
			handler.cache.Set("elb", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) ELBRequestsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("elb_requests")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetELBRequests(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "elasticloadbalancing:DescribeLoadBalancers is missing")
		} else {
			handler.cache.Set("elb_requests", response)
			respondWithJSON(w, 200, response)
		}
	}
}
