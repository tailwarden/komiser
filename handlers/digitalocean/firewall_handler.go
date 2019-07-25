package digitalocean

import (
	"net/http"
)

func (handler *DigitalOceanHandler) DescribeFirewallsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.firewalls.list")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeFirewalls(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.firewalls.list", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *DigitalOceanHandler) DescribeUnsecureFirewallsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("digitalocean.firewalls.unsecure")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.digitalocean.DescribeUnsecureFirewalls(handler.client)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Not enough permissions to access DigitalOcean API")
		} else {
			handler.cache.Set("digitalocean.firewalls.unsecure", response)
			respondWithJSON(w, 200, response)
		}
	}
}
