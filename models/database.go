package models

type DatabaseConfig struct {
	Type     string `json:"type"`
	Hostname string `json:"hostname"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	File     string `json:"file"`
}
