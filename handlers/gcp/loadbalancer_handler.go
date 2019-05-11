package aws

import (
	"net/http"
)

func (handler *GCPHandler) LoadBalancersRequestsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_lb_requests")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetLoadBalancerRequests()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "monitoring:MonitoringReadScope is missing")
		} else {
			handler.cache.Set("gcp_lb_requests", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) LoadBalancersTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_lb_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetTotalLoadBalancers()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("gcp_lb_total", response)
			respondWithJSON(w, 200, response)
		}
	}
}
