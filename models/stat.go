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

type InputCostBreakdown struct {
	Group       string   `json:"group"`
	Granularity string   `json:"granularity"`
	Start       string   `json:"start"`
	End         string   `json:"end"`
	Exclude     []string `json:"exclude"`
}

type OutputCostBreakdownRaw struct {
	Provider string  `json:"provider"`
	Account  string  `json:"account"`
	Region   string  `json:"region"`
	Service  string  `json:"service"`
	Total    float64 `json:"total"`
	Period   string  `json:"period"`
}

type OutputCostBreakdown struct {
	Date       string      `json:"date"`
	Datapoints []Datapoint `json:"datapoints"`
}

type Datapoint struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type ViewStat struct {
	Resources int     `json:"resources"`
	Costs     float64 `json:"costs"`
}

type OutputCostByField struct {
	Label string  `bun:"label"`
	Total float64 `bun:"total"`
}
