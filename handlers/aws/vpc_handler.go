package aws

import (
	"net/http"
)

func (handler *AWSHandler) VPCHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("vpc")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeVPCsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeVpcs is missing")
		} else {
			handler.cache.Set("vpc", response)
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
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeNetworkAcls is missing")
		} else {
			handler.cache.Set("acl", response)
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
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeSecurityGroups is missing")
		} else {
			handler.cache.Set("sg", response)
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
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeNatGateways is missing")
		} else {
			handler.cache.Set("nat", response)
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
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeElasticIPs is missing")
		} else {
			handler.cache.Set("eip", response)
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
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeInternetGateways is missing")
		} else {
			handler.cache.Set("igw", response)
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
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeRouteTables is missing")
		} else {
			handler.cache.Set("rt", response)
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
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeKeyPairs is missing")
		} else {
			handler.cache.Set("kp", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) GetNatGatewayTrafficHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("nat_traffic")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetNatGatewayTraffic(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set("nat_traffic", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) DescribeSubnetsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("subnets")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeSubnets(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeSubnets is missing")
		} else {
			handler.cache.Set("subnets", response)
			respondWithJSON(w, 200, response)
		}
	}
}
