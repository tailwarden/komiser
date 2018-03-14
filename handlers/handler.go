package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	. "github.com/mlabouardy/komiser/backend"
	cache "github.com/patrickmn/go-cache"
)

type AWSHandler struct {
	cfg   aws.Config
	cache *cache.Cache
	aws   AWS
}

func NewAWSHandler(cfg aws.Config, cache *cache.Cache) *AWSHandler {
	awsHandler := AWSHandler{
		cfg:   cfg,
		cache: cache,
		aws:   AWS{},
	}
	return &awsHandler
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (handler *AWSHandler) EC2RegionHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ec2_region")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeInstancesPerRegion(handler.cfg)
		handler.cache.Set("ec2_region", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) EC2FamilyHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ec2_family")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeInstancesPerFamily(handler.cfg)
		handler.cache.Set("ec2_family", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) EC2StateHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ec2_state")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeInstancesPerState(handler.cfg)
		handler.cache.Set("ec2_state", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) EBSSizeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ebs_size")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeVolumesTotalSize(handler.cfg)
		handler.cache.Set("ebs_size", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) EBSFamilyHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ebs_family")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeVolumesPerFamily(handler.cfg)
		handler.cache.Set("ebs_family", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) EBSStateHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ebs_state")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeVolumesPerState(handler.cfg)
		handler.cache.Set("ebs_state", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) VPCTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("vpc_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeVPCsTotal(handler.cfg)
		handler.cache.Set("vpc_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) ACLTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("acl_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeACLsTotal(handler.cfg)
		handler.cache.Set("acl_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) SecurityGroupTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("sg_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeSecurityGroupsTotal(handler.cfg)
		handler.cache.Set("sg_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) NatGatewayTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("nat_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeNatGatewaysTotal(handler.cfg)
		handler.cache.Set("nat_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) ElasticIPTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("eip_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeElasticIPsTotal(handler.cfg)
		handler.cache.Set("eip_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) InternetGatewayTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("igw_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeInternetGatewaysTotal(handler.cfg)
		handler.cache.Set("igw_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) RouteTableTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("rt_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeRouteTablesTotal(handler.cfg)
		handler.cache.Set("rt_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) KeyPairTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("kp_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeKeyPairsTotal(handler.cfg)
		handler.cache.Set("kp_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) AutoScalingGroupTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("asg_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeAutoScalingGroupsTotal(handler.cfg)
		handler.cache.Set("asg_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) ElasticLoadBalancerFamilyHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("elb_family")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeElasticLoadBalancerPerFamily(handler.cfg)
		handler.cache.Set("elb_family", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) S3TotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("s3_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeS3BucketsTotal(handler.cfg)
		handler.cache.Set("s3_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) CostAndUsageHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cost_usage")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeCostAndUsage(handler.cfg)
		handler.cache.Set("cost_usage", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) LambdaPerRuntimeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("lambda_runtime")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeLambdaFunctionsPerRuntime(handler.cfg)
		handler.cache.Set("lambda_runtime", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) RDSInstancePerEngineHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("rds_engine")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeRDSInstancesPerEngine(handler.cfg)
		handler.cache.Set("rds_engine", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) DynamoDBTableTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("dynamodb_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeDynamoDBTablesTotal(handler.cfg)
		handler.cache.Set("dynamodb_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) DynamoDBProvisionedThroughputHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("dynamodb_throughput")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeDynamoDBTablesProvisionedThroughput(handler.cfg)
		handler.cache.Set("dynamodb_throughput", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) SnapshotTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("snapshot_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeSnapshotsTotal(handler.cfg)
		handler.cache.Set("snapshot_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) SnapshotSizeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("snapshot_size")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeSnapshotsSize(handler.cfg)
		handler.cache.Set("snapshot_size", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) SQSTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("sqs_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeQueuesTotal(handler.cfg)
		handler.cache.Set("sqs_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) TopicsTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("sns_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeSNSTopicsTotal(handler.cfg)
		handler.cache.Set("sns_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) HostedZoneTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("hosted_zone_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeHostedZonesTotal(handler.cfg)
		handler.cache.Set("hosted_zone_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) IAMRolesTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("role_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeIAMRolesTotal(handler.cfg)
		handler.cache.Set("role_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) IAMGroupsTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("group_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeIAMGroupsTotal(handler.cfg)
		handler.cache.Set("group_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) IAMPoliciesTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("policy_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeIAMPoliciesTotal(handler.cfg)
		handler.cache.Set("policy_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) IAMUsersTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("user_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeIAMUsersTotal(handler.cfg)
		handler.cache.Set("user_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) CloudWatchAlarmsStateHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cloudwatch_alarm")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeCloudWatchAlarmsPerState(handler.cfg)
		handler.cache.Set("cloudwatch_alarm", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}

func (handler *AWSHandler) CloudFrontDistributionsTotalHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cloudfront_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response := handler.aws.DescribeCloudFrontDistributionsTotal(handler.cfg)
		handler.cache.Set("cloudfront_total", response, cache.DefaultExpiration)
		respondWithJSON(w, 200, response)
	}
}
