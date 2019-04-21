package aws

import "time"

type Bill struct {
	Total   float64 `json:"total"`
	History []Cost  `json:"history"`
}

type Cost struct {
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Unit   string    `json:"unit"`
	Groups []Group   `json:"groups"`
}

type Group struct {
	Key    string  `json:"key"`
	Amount float64 `json:"amount"`
}
