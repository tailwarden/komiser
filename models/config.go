package models

type Config struct {
	AWS          []AWSConfig          `toml:"aws"`
	DigitalOcean []DigitalOceanConfig `toml:"digitalocean"`
	Oci          []OciConfig          `toml:"oci"`
	Civo         []CivoConfig         `toml:"civo"`
	Kubernetes   []KubernetesConfig   `toml:"k8s"`
	Linode       []LinodeConfig       `toml:"linode"`
	Tencent      []TencentConfig      `toml:"tencent"`
	Postgres     PostgresConfig       `toml:"postgres,omitempty"`
	SQLite       SQLiteConfig         `toml:"sqlite"`
}

type AWSConfig struct {
	Name    string `toml:"name"`
	Profile string `toml:"profile"`
	Source  string `toml:"source"`
	Path    string `toml:"path,omitempty"`
}

type DigitalOceanConfig struct {
	Name  string `toml:"name"`
	Token string `toml:"token"`
}

type KubernetesConfig struct {
	Name     string   `toml:"name"`
	Path     string   `toml:"path"`
	Contexts []string `toml:"contexts"`
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
