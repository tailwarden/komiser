package models

type View struct {
	Id      int64    `json:"id" bun:"id,pk,autoincrement"`
	Name    string   `json:"name"`
	Filters []Filter `json:"filters"`
	Exclude []int64  `json:"exclude"`
}
