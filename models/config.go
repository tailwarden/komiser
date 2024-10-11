package models

type Config struct {
	AWS          []AWSConfig          `toml:"aws"`
	Azure        []AzureConfig        `toml:"azure"`
	DigitalOcean []DigitalOceanConfig `toml:"digitalocean"`
	Oci          []OciConfig          `toml:"oci"`
	Civo         []CivoConfig         `toml:"civo"`
	Kubernetes   []KubernetesConfig   `toml:"k8s"`
	Linode       []LinodeConfig       `toml:"linode"`
	Tencent      []TencentConfig      `toml:"tencent"`
	Scaleway     []ScalewayConfig     `toml:"scaleway"`
	MongoDBAtlas []MongoDBAtlasConfig `toml:"mongodbatlas"`
	GCP          []GCPConfig          `toml:"gcp"`
	OVH          []OVHConfig          `toml:"ovh"`
	Postgres     PostgresConfig       `toml:"postgres,omitempty"`
	SQLite       SQLiteConfig         `toml:"sqlite"`
	Slack        SlackConfig          `toml:"slack"`
}

type AWSConfig struct {
	Name    string `toml:"name"`
	Profile string `toml:"profile"`
	Source  string `toml:"source"`
	Path    string `toml:"path,omitempty"`
}

type AzureConfig struct {
	Name           string `toml:"name"`
	TenantId       string `toml:"tenantId"`
	ClientId       string `toml:"clientId"`
	ClientSecret   string `toml:"clientSecret"`
	SubscriptionId string `toml:"subscriptionId"`
}

type DigitalOceanConfig struct {
	Name  string `toml:"name"`
	Token string `toml:"token"`
}

type ScalewayConfig struct {
	Name           string `toml:"name"`
	OrganizationId string `toml:"organizationId"`
	AccessKey      string `toml:"accessKey"`
	SecretKey      string `toml:"secretKey"`
}

type KubernetesConfig struct {
	Name            string   `toml:"name"`
	Path            string   `toml:"path"`
	Contexts        []string `toml:"contexts"`
	OpencostBaseUrl string   `toml:"opencostBaseUrl"`
}

type LinodeConfig struct {
	Name  string `toml:"name"`
	Token string `toml:"token"`
}

type CivoConfig struct {
	Name  string `toml:"name"`
	Token string `toml:"token"`
}

type PostgresConfig struct {
	URI string `toml:"uri,omitempty"`
}

type SQLiteConfig struct {
	File string `toml:"file"`
}

type OciConfig struct {
	Name    string `toml:"name"`
	Profile string `toml:"profile"`
	Source  string `toml:"source"`
}

// TencentConfig holds the configuration for Tencent cloud.
type TencentConfig struct {
	Name      string `toml:"name"`
	SecretID  string `toml:"secret_id"`
	SecretKey string `toml:"secret_key"`
}

type MongoDBAtlasConfig struct {
	Name           string `toml:"name"`
	OrganizationID string `toml:"organizationId"`
	PublicApiKey   string `toml:"publicApiKey"`
	PrivateApiKey  string `toml:"privateApiKey"`
}

type GCPConfig struct {
	Name                  string `toml:"name"`
	ServiceAccountKeyPath string `toml:"serviceAccountKeyPath"`
}

type OVHConfig struct {
	Name              string `toml:"name"`
	Endpoint          string `toml:"endpoint"`
	ApplicationKey    string `toml:"application_key"`
	ApplicationSecret string `toml:"application_secret"`
	ConsumerKey       string `toml:"consumer_key"`
}

type SlackConfig struct {
	Webhook   string `toml:"webhook"`
	Reporting bool   `toml:"reporting"`
	Host      string `toml:"host"`
}
