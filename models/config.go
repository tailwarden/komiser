package models

type Config struct {
	AWS          []AWSConfig          `toml:"aws"`
	DigitalOcean []DigitalOceanConfig `toml:"digitalocean"`
	Oci          []OciConfig          `toml:"oci"`
	Civo         []CivoConfig         `toml:"civo"`
	Kubernetes   []KubernetesConfig   `toml:"k8s"`
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
