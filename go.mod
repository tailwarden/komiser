module github.com/tailwarden/komiser

go 1.21

require (
	cloud.google.com/go/bigquery v1.49.0
	cloud.google.com/go/compute v1.19.0
	cloud.google.com/go/container v1.16.0
	cloud.google.com/go/kms v1.10.0
	cloud.google.com/go/monitoring v1.13.0
	cloud.google.com/go/redis v1.11.0
	cloud.google.com/go/storage v1.30.0
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.3.1
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.2.1
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4 v4.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork v1.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage v1.2.0
	github.com/BurntSushi/toml v1.2.1
	github.com/aws/aws-sdk-go-v2 v1.21.0
	github.com/aws/aws-sdk-go-v2/config v1.18.27
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.16.0
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.27.2
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.20.7
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.23.1
	github.com/aws/aws-sdk-go-v2/service/configservice v1.31.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.17.3
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.50.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.17.20
	github.com/aws/aws-sdk-go-v2/service/ecs v1.18.26
	github.com/aws/aws-sdk-go-v2/service/efs v1.19.0
	github.com/aws/aws-sdk-go-v2/service/eks v1.21.4
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.26.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.19.4
	github.com/aws/aws-sdk-go-v2/service/iam v1.18.9
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.17.12
	github.com/aws/aws-sdk-go-v2/service/kms v1.18.18
	github.com/aws/aws-sdk-go-v2/service/lambda v1.23.4
	github.com/aws/aws-sdk-go-v2/service/pricing v1.20.0
	github.com/aws/aws-sdk-go-v2/service/rds v1.30.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.29.5
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.1
	github.com/aws/aws-sdk-go-v2/service/sns v1.18.3
	github.com/aws/aws-sdk-go-v2/service/sqs v1.19.12
	github.com/aws/aws-sdk-go-v2/service/sts v1.19.2
	github.com/civo/civogo v0.3.24
	github.com/digitalocean/godo v1.97.0
	github.com/elazarl/go-bindata-assetfs v1.0.1
	github.com/getsentry/sentry-go v0.18.0
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-gonic/gin v1.9.1
	github.com/go-co-op/gocron v1.18.0
	github.com/golang/protobuf v1.5.3
	github.com/hashicorp/go-version v1.6.0
	github.com/linode/linodego v1.12.0
	github.com/mongodb-forks/digest v1.0.4
	github.com/oracle/oci-go-sdk v24.3.0+incompatible
	github.com/rs/xid v1.4.0
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.13
	github.com/segmentio/analytics-go v3.1.0+incompatible
	github.com/sirupsen/logrus v1.9.0
	github.com/siruspen/logrus v1.7.1
	github.com/slack-go/slack v0.12.1
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.3
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.582
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm v1.0.582
	github.com/uptrace/bun v1.1.8
	github.com/uptrace/bun/dialect/pgdialect v1.1.8
	github.com/uptrace/bun/dialect/sqlitedialect v1.1.8
	github.com/uptrace/bun/driver/pgdriver v1.1.8
	github.com/uptrace/bun/driver/sqliteshim v1.1.8
	go.mongodb.org/atlas v0.23.1
	golang.org/x/oauth2 v0.6.0
	golang.org/x/text v0.9.0
	google.golang.org/api v0.114.0
	k8s.io/apimachinery v0.26.1
	k8s.io/client-go v0.26.1
)

require (
	cloud.google.com/go/longrunning v0.4.1 // indirect
	github.com/apache/arrow/go/v11 v11.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.12 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
)

require (
	cloud.google.com/go v0.110.0 // indirect
	cloud.google.com/go/certificatemanager v1.6.0
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v0.13.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.1.2 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v0.8.1 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/apache/thrift v0.16.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.13.26 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.41 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.35 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.35 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.20.9
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.28 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.15.4
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.18.2
	github.com/aws/aws-sdk-go-v2/service/ssm v1.36.2
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.12 // indirect
	github.com/aws/smithy-go v1.14.2 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/flatbuffers v2.0.8+incompatible // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.7.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/imdario/mergo v0.3.6 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/klauspost/asmfmt v1.3.2 // indirect
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/minio/asm2plan9s v0.0.0-20200509001527-cdd76441f9d8 // indirect
	github.com/minio/c2goasm v0.0.0-20190812172519-36a3d3bbc4f3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/openlyinc/pointy v1.2.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/pkg/browser v0.0.0-20210115035449-ce105d075bb4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/segmentio/backo-go v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xtgo/uuid v0.0.0-20140804021211-a0b114877d4c // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/term v0.8.0 // indirect
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af // indirect
	golang.org/x/tools v0.6.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230330154414-c0448cd141ea // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.26.1 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
	k8s.io/kube-openapi v0.0.0-20221012153701-172d655c2280 // indirect
	k8s.io/utils v0.0.0-20221107191617-1a15be271d1d // indirect
	lukechampine.com/uint128 v1.2.0 // indirect
	mellium.im/sasl v0.3.0 // indirect
	modernc.org/cc/v3 v3.36.3 // indirect
	modernc.org/ccgo/v3 v3.16.9 // indirect
	modernc.org/libc v1.17.1 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.2.1 // indirect
	modernc.org/opt v0.1.3 // indirect
	modernc.org/sqlite v1.18.1 // indirect
	modernc.org/strutil v1.1.3 // indirect
	modernc.org/token v1.0.1 // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
