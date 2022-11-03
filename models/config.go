package models

type ConfigFile struct {
	AWS          []AWSConfig
	DigitalOcean []DigitalOceanConfig
	Postgres     PostgresConfig
}

type AWSConfig struct {
	Name    string
	Profile string
	Source  string
}

type DigitalOceanConfig struct {
	Name  string
	Token string
}

type PostgresConfig struct {
	URI string
}
