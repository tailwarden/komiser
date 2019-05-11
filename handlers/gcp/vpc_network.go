package aws

import (
	"fmt"
	"net/http"
)

func (handler *GCPHandler) VpcNetworksHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("vpc_networks")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetVpcNetworks()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("vpc_networks", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) VpcFirewallsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("vpc_firewalls")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetNetworkFirewalls()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("vpc_firewalls", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) VpcRoutersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("vpc_routers")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetNetworkRouters()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("vpc_routers", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) VpcSubnetsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("vpc_subnets")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetSubnetsNumber()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("vpc_subnets", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) VpcExternalAddressesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("vpc_addresses")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetExternalAddresses()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("vpc_addresses", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) VpnTunnelsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("vpn_tunnels")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetVpnTunnels()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("vpn_tunnels", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) SSLCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ssl_certificates")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetSSLCertificates()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("ssl_certificates", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) SSLPoliciesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ssl_policies")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetSSLPolicies()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("ssl_policies", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) SecurityPoliciesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("security_policies")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetSecurityPolicies()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("security_policies", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) NatGatewaysHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("nat_gateways")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetNatGateways()
		if err != nil {
			fmt.Println(err)
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("nat_gateways", response)
			respondWithJSON(w, 200, response)
		}
	}
}
