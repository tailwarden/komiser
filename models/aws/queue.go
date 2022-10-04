package aws

type Queue struct {
	Name   string   `json:"name"`
	Region string   `json:"region"`
	ID     string   `json:"id"`
	Tags   []string `json:"tags"`
}
