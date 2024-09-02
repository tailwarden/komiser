module github.com/tailwarden/komiser

go 1.21

require (
	cloud.google.com/go/alloydb v1.10.1
	cloud.google.com/go/appengine v1.8.5
	cloud.google.com/go/artifactregistry v1.14.7
	cloud.google.com/go/bigquery v1.59.1
	cloud.google.com/go/compute v1.24.0
	cloud.google.com/go/container v1.31.0
	cloud.google.com/go/firestore v1.15.0
	cloud.google.com/go/kms v1.15.7
	cloud.google.com/go/monitoring v1.18.0
	cloud.google.com/go/redis v1.14.2
	cloud.google.com/go/storage v1.38.0
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.11.1
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.6.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4 v4.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement v1.1.1
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork v1.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v1.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage v1.2.0
	github.com/BurntSushi/toml v1.2.1
	github.com/aws/aws-sdk-go-v2 v1.25.2
	github.com/aws/aws-sdk-go-v2/config v1.25.3
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.20.2
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.35.1
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.30.2
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.38.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.30.2
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.18.2
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.20.3
	github.com/aws/aws-sdk-go-v2/service/configservice v1.41.2
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.32.4
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.25.2
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.136.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.23.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.33.1
	github.com/aws/aws-sdk-go-v2/service/efs v1.23.2
	github.com/aws/aws-sdk-go-v2/service/eks v1.33.1
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.32.2
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.24.2
	github.com/aws/aws-sdk-go-v2/service/firehose v1.28.1
	github.com/aws/aws-sdk-go-v2/service/iam v1.27.2
	github.com/aws/aws-sdk-go-v2/service/kafka v1.30.1
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.22.2
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.23.1
	github.com/aws/aws-sdk-go-v2/service/kms v1.26.2
	github.com/aws/aws-sdk-go-v2/service/lambda v1.48.0
	github.com/aws/aws-sdk-go-v2/service/pricing v1.23.2
	github.com/aws/aws-sdk-go-v2/service/rds v1.63.0
	github.com/aws/aws-sdk-go-v2/service/redshift v1.37.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.43.0
	github.com/aws/aws-sdk-go-v2/service/sns v1.25.2
	github.com/aws/aws-sdk-go-v2/service/sqs v1.28.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.25.3
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.47.0
	github.com/civo/civogo v0.3.24
	github.com/digitalocean/godo v1.97.0
	github.com/elazarl/go-bindata-assetfs v1.0.1
	github.com/getsentry/sentry-go v0.18.0
	github.com/gin-contrib/cors v1.6.0
	github.com/gin-gonic/gin v1.9.1
	github.com/go-co-op/gocron v1.18.0
	github.com/golang/protobuf v1.5.3
	github.com/hashicorp/go-version v1.6.0
	github.com/linode/linodego v1.12.0
	github.com/mongodb-forks/digest v1.0.4
	github.com/oracle/oci-go-sdk v24.3.0+incompatible
	github.com/ovh/go-ovh v1.4.3
	github.com/rs/xid v1.4.0
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.13
	github.com/segmentio/analytics-go v3.1.0+incompatible
	github.com/sirupsen/logrus v1.9.3
	github.com/siruspen/logrus v1.7.1
	github.com/slack-go/slack v0.12.1
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.9.0
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.582
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm v1.0.582
	github.com/uptrace/bun v1.1.8
	github.com/uptrace/bun/dialect/pgdialect v1.1.8
	github.com/uptrace/bun/dialect/sqlitedialect v1.1.8
	github.com/uptrace/bun/driver/pgdriver v1.1.8
	github.com/uptrace/bun/driver/sqliteshim v1.1.8
	go.mongodb.org/atlas v0.23.1
	golang.org/x/oauth2 v0.17.0
	golang.org/x/text v0.16.0
	google.golang.org/api v0.169.0
	k8s.io/apimachinery v0.26.1
	k8s.io/client-go v0.26.1
)

require (
	cloud.google.com/go/longrunning v0.5.5 // indirect
	github.com/apache/arrow/go/v14 v14.0.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.20.0 // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.49.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.49.0 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240311132316-a219d84964c2 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240311132316-a219d84964c2 // indirect
)

require (
	cloud.google.com/go v0.112.1 // indirect
	cloud.google.com/go/certificatemanager v1.7.5
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v1.1.6 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.8.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.2.2 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.1 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.16.2 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.2.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.27.1
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.25.2
	github.com/aws/aws-sdk-go-v2/service/datasync v1.36.0
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.2.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.16.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.35.0
	github.com/aws/aws-sdk-go-v2/service/neptune v1.29.0
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.25.2
	github.com/aws/aws-sdk-go-v2/service/route53 v1.38.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.28.0
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.24.2
	github.com/aws/aws-sdk-go-v2/service/ssm v1.43.0
	github.com/aws/aws-sdk-go-v2/service/sso v1.17.2 // indirect
	github.com/aws/smithy-go v1.20.1 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/bytedance/sonic v1.11.2 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.19.0 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/flatbuffers v23.5.26+incompatible // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.2 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/imdario/mergo v0.3.6 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/openlyinc/pointy v1.2.0 // indirect
	github.com/pelletier/go-toml/v2 v2.1.1 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/segmentio/backo-go v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xtgo/uuid v0.0.0-20140804021211-a0b114877d4c // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/arch v0.7.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/term v0.21.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20240213162025-012b6fc9bca9 // indirect
	google.golang.org/grpc v1.62.1 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.26.1 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
	k8s.io/kube-openapi v0.0.0-20221012153701-172d655c2280 // indirect
	k8s.io/utils v0.0.0-20221107191617-1a15be271d1d // indirect
	lukechampine.com/uint128 v1.3.0 // indirect
	mellium.im/sasl v0.3.0 // indirect
	modernc.org/cc/v3 v3.40.0 // indirect
	modernc.org/ccgo/v3 v3.16.13 // indirect
	modernc.org/libc v1.22.4 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.5.0 // indirect
	modernc.org/opt v0.1.3 // indirect
	modernc.org/sqlite v1.21.2 // indirect
	modernc.org/strutil v1.1.3 // indirect
	modernc.org/token v1.1.0 // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
