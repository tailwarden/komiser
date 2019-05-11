package gcp

type Cost struct {
	Date string  `json:"date"`
	Cost float64 `json:"cost"`
}

type CostPerService struct {
	Date   string  `json:"date"`
	Unit   string  `json:"unit"`
	Groups []Group `json:"groups"`
}

type Group struct {
	Service string  `json:"service"`
	Unit    string  `json:"unit"`
	Cost    float64 `json:"cost"`
}
