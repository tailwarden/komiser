package aws

import (
	"net/http"
)

func (handler *AWSHandler) VPCHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_vpc")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeVPCsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeVpcs is missing")
		} else {
			handler.cache.Set("aws_vpc", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) ACLHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_acl")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeACLsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeNetworkAcls is missing")
		} else {
			handler.cache.Set("aws_acl", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) SecurityGroupHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_sg")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeSecurityGroupsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeSecurityGroups is missing")
		} else {
			handler.cache.Set("aws_sg", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) NatGatewayHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_nat")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeNatGatewaysTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeNatGateways is missing")
		} else {
			handler.cache.Set("aws_nat", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) ElasticIPHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_eip")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeElasticIPsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeElasticIPs is missing")
		} else {
			handler.cache.Set("aws_eip", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) InternetGatewayHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_igw")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeInternetGatewaysTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeInternetGateways is missing")
		} else {
			handler.cache.Set("aws_igw", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) RouteTableHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_rt")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeRouteTablesTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeRouteTables is missing")
		} else {
			handler.cache.Set("aws_rt", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) KeyPairHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_kp")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeKeyPairsTotal(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeKeyPairs is missing")
		} else {
			handler.cache.Set("aws_kp", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) GetNatGatewayTrafficHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_nat_traffic")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetNatGatewayTraffic(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudwatch:GetMetricStatistics is missing")
		} else {
			handler.cache.Set("aws_nat_traffic", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) DescribeSubnetsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_subnets")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeSubnets(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ec2:DescribeSubnets is missing")
		} else {
			handler.cache.Set("aws_subnets", response)
			respondWithJSON(w, 200, response)
		}
	}
}
