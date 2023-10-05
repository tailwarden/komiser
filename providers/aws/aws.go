package aws

import (
	"context"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/aws/apigateway"
	"github.com/tailwarden/komiser/providers/aws/cloudfront"
	"github.com/tailwarden/komiser/providers/aws/cloudwatch"
	"github.com/tailwarden/komiser/providers/aws/dynamodb"
	"github.com/tailwarden/komiser/providers/aws/ec2"
	"github.com/tailwarden/komiser/providers/aws/ecr"
	"github.com/tailwarden/komiser/providers/aws/ecs"
	"github.com/tailwarden/komiser/providers/aws/efs"
	"github.com/tailwarden/komiser/providers/aws/eks"
	"github.com/tailwarden/komiser/providers/aws/elasticache"
	"github.com/tailwarden/komiser/providers/aws/elb"
	"github.com/tailwarden/komiser/providers/aws/iam"
	"github.com/tailwarden/komiser/providers/aws/kinesis"
	"github.com/tailwarden/komiser/providers/aws/kms"
	"github.com/tailwarden/komiser/providers/aws/lambda"
	"github.com/tailwarden/komiser/providers/aws/opensearch"
	"github.com/tailwarden/komiser/providers/aws/rds"
	"github.com/tailwarden/komiser/providers/aws/redshift"
	"github.com/tailwarden/komiser/providers/aws/s3"
	"github.com/tailwarden/komiser/providers/aws/servicecatalog"
	"github.com/tailwarden/komiser/providers/aws/sns"
	"github.com/tailwarden/komiser/providers/aws/sqs"
	"github.com/tailwarden/komiser/providers/aws/systemsmanager"
	"github.com/tailwarden/komiser/utils"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		ec2.Instances,
		ec2.ElasticIps,
		lambda.Functions,
		lambda.EventSourceMappings,
		ec2.Acls,
		ec2.Subnets,
		ec2.SecurityGroups,
		ec2.AutoScalingGroups,
		ec2.InternetGateways,
		iam.Roles,
		iam.InstanceProfiles,
		iam.OIDCProviders,
		iam.SamlProviders,
		iam.Groups,
		iam.Policies,
		iam.Users,
		sqs.Queues,
		s3.Buckets,
		ec2.Instances,
		eks.KubernetesClusters,
		cloudfront.Distributions,
		dynamodb.Tables,
		ecs.Clusters,
		ecs.TaskDefinitions,
		ecs.ContainerInstances,
		ecr.Repositories,
		sns.Topics,
		ec2.Vpcs,
		ec2.Volumes,
		kms.Keys,
		rds.Clusters,
		rds.Instances,
		rds.Proxies,
		rds.Snapshots,
		rds.ClusterSnapshots,
		rds.ProxyEndpoints,
		rds.AutoBackups,
		elb.LoadBalancers,
		elb.TargetGroups,
		efs.ElasticFileStorage,
		apigateway.Apis,
		elasticache.Clusters,
		cloudwatch.Alarms,
		ec2.NetworkInterfaces,
		cloudwatch.Dashboards,
		ec2.ElasticIps,
		cloudwatch.LogGroups,
		cloudwatch.MetricStreams,
		ec2.Snapshots,
		opensearch.ServiceDomains,
		servicecatalog.Products,
		ec2.SpotInstanceRequests,
		ec2.KeyPairs,
		ec2.PlacementGroups,
		systemsmanager.MaintenanceWindows,
		ec2.VpcEndpoints,
		ec2.VpcPeeringConnections,
		kinesis.Streams,
		redshift.EventSubscriptions,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, regions []string, db *bun.DB, telemetry bool, analytics utils.Analytics) {
	listOfSupportedRegions := getRegions()
	if len(regions) > 0 {
		log.Infof("Komiser will fetch resources from the following regions: %s", strings.Join(regions, ","))
		listOfSupportedRegions = regions
	}

	for _, region := range listOfSupportedRegions {
		client.AWSClient.Region = region
		for _, fetchResources := range listOfSupportedServices() {
			resources, err := fetchResources(ctx, client)
			if err != nil {
				log.Warnf("[%s][AWS] %s", client.Name, err)
			} else {
				for _, resource := range resources {
				_, err = db.NewInsert().Model(&resource).On("CONFLICT (resource_id) DO UPDATE").Set("cost = EXCLUDED.cost, relations=EXCLUDED.relations").Exec(context.Background())
					if err != nil {
						log.WithError(err).Errorf("db trigger failed")
					}
				}
				if telemetry {
					analytics.TrackEvent("discovered_resources", map[string]interface{}{
						"provider":  "AWS",
						"resources": len(resources),
					})
				}
			}
		}
	}
}

func getRegions() []string {
	return []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2", "ca-central-1", "eu-north-1", "eu-west-1", "eu-west-2", "eu-west-3", "eu-central-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-southeast-1", "ap-southeast-2", "ap-south-1", "sa-east-1"}
}
