module github.com/mlabouardy/komiser

go 1.17

require (
	cloud.google.com/go/bigquery v1.36.0
	github.com/Azure/azure-sdk-for-go v66.0.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.28
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.11
	github.com/aws/aws-sdk-go v1.44.63
	github.com/aws/aws-sdk-go-v2 v1.16.7
	github.com/aws/aws-sdk-go-v2/config v1.15.14
	github.com/aws/aws-sdk-go-v2/service/acm v1.14.8
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.15.10
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.23.6
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.18.4
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.16.4
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.19.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.15.10
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.19.2
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.13.8
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.15.9
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.50.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.18.11
	github.com/aws/aws-sdk-go-v2/service/eks v1.21.4
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.22.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.14.8
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.18.8
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.15.9
	github.com/aws/aws-sdk-go-v2/service/glacier v1.13.8
	github.com/aws/aws-sdk-go-v2/service/glue v1.28.1
	github.com/aws/aws-sdk-go-v2/service/iam v1.18.9
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.15.9
	github.com/aws/aws-sdk-go-v2/service/kms v1.18.0
	github.com/aws/aws-sdk-go-v2/service/lambda v1.23.4
	github.com/aws/aws-sdk-go-v2/service/mq v1.13.4
	github.com/aws/aws-sdk-go-v2/service/organizations v1.16.4
	github.com/aws/aws-sdk-go-v2/service/rds v1.23.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.26.0
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.13.9
	github.com/aws/aws-sdk-go-v2/service/route53 v1.21.3
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.1
	github.com/aws/aws-sdk-go-v2/service/sns v1.17.9
	github.com/aws/aws-sdk-go-v2/service/sqs v1.19.0
	github.com/aws/aws-sdk-go-v2/service/support v1.13.8
	github.com/aws/aws-sdk-go-v2/service/swf v1.13.8
	github.com/digitalocean/godo v1.81.0
	github.com/elazarl/go-bindata-assetfs v1.0.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/ovh/go-ovh v1.1.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron v1.2.0
	github.com/rs/cors v1.8.2
	github.com/slack-go/slack v0.11.2
	github.com/urfave/cli v1.22.9
	golang.org/x/oauth2 v0.0.0-20220722155238-128564f6959c
	google.golang.org/api v0.89.0
)

require (
	cloud.google.com/go v0.103.0 // indirect
	cloud.google.com/go/compute v1.7.0 // indirect
	cloud.google.com/go/iam v0.3.0 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.21 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.6 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.3 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.12.9 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.9 // indirect
	github.com/aws/smithy-go v1.12.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.1.0 // indirect
	github.com/googleapis/gax-go/v2 v2.4.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.18.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220726230323-06994584191e // indirect
	golang.org/x/sys v0.0.0-20220727055044-e65921a090b8 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220725144611-272f38e5d71b // indirect
	google.golang.org/grpc v1.48.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/ini.v1 v1.66.6 // indirect
)
