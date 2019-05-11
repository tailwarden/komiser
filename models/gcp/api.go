package gcp

type API struct {
	Namespace string `json:"namespace"`
	Title     string `json:"title"`
	Enabled   bool   `json:"enabled"`
}
