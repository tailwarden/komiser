package models

type Link struct {
	ResourceID string
	Type string
	Relation string
}

type OutputRelationResponse struct {
	ResourceID string `json:"resourceId" bun:"resource_id,unique"`
	Type string `json:"service" bun:"service"`
	Link []Link `json:"relations" bun:"relations"`
}