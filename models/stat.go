package models

type InputResources struct {
	Filter  string   `json:"filter"`
	Exclude []string `json:"exclude"`
}

type OutputResources struct {
	Label string `json:"label"`
	Total int    `json:"total"`
}
