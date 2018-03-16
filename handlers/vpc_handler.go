package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) VPCHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("vpc")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeVPCsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("vpc", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) ACLHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("acl")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeACLsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("acl", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) SecurityGroupHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("sg")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeSecurityGroupsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("sg", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) NatGatewayHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("nat")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeNatGatewaysTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("nat", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) ElasticIPHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("eip")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeElasticIPsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("eip", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) InternetGatewayHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("igw")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeInternetGatewaysTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("igw", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) RouteTableHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("rt")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeRouteTablesTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("rt", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) KeyPairHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("kp")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeKeyPairsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "You dont have the right permission")
		} else {
			handler.cache.Set("kp", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
