package models

type Alert struct {
	Id       int64   `json:"id" bun:"id,pk,autoincrement"`
	Name     string  `json:"name"`
	ViewId   string  `json:"viewId" bun:"view_id"`
	Type     string  `json:"type"`
	Budget   float64 `json:"budget"`
	Usage    int     `json:"usage"`
	IsSlack  bool    `json:"isSlack" bun:"is_slack"`
	Endpoint string  `json:"endpoint"`
	Secret   string  `json:"secret"`
}

type Endpoint struct {
	Url string `json:"url"`
}

type CustomWebhookPayload struct {
	Komiser   string  `json:"komiser"`
	View      string  `json:"view"`
	Message   string  `json:"message"`
	Data      float64 `json:"data"`
	Timestamp int64   `json:"timestamp"`
}
