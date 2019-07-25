package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) LoadBalancersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.loadbalancers")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeLoadBalancers(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.loadbalancers", response)
			respondWithJSON(w, 200, response)
		}
	}
}
