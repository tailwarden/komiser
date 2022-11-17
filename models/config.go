package models

type Config struct {
	AWS          []AWSConfig
	DigitalOcean []DigitalOceanConfig
	Oci          []OciConfig
	Postgres     PostgresConfig
}

type AWSConfig struct {
	Name    string
	Profile string
	Source  string
	Path    string
}

type DigitalOceanConfig struct {
	Name  string
	Token string
}

type PostgresConfig struct {
	URI string
}

type OciConfig struct {
	Name    string
	Profile string
	Source  string
}
