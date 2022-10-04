package aws

type SecurityGroup struct {
	Name   string   `json:"name"`
	ID     string   `json:"id"`
	Region string   `json:"region"`
	Tags   []string `json:"tags"`
}
