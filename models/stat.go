package models

type InputResources struct {
	Filter  string   `json:"filter"`
	Exclude []string `json:"exclude"`
}

type OutputResources struct {
	Label string `json:"label"`
	Total int    `json:"total"`
}

type OutputLocations struct {
	Name      string `json:"name"`
	Label     string `json:"label"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Resources int    `json:"resources"`
}
