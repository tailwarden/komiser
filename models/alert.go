package models

type Alert struct {
	Id     int64   `json:"id" bun:"id,pk,autoincrement"`
	Name   string  `json:"name"`
	ViewId string  `json:"viewId" bun:"view_id"`
	Type   string  `json:"type"`
	Budget float64 `json:"budget"`
	Usage  int     `json:"usage"`
}
