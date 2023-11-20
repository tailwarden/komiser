package models

type Link struct {
	ResourceID string `json:"resourceId"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Relation   string `json:"relation"`
}

type OutputRelationResponse struct {
	ResourceID string `json:"resourceId" bun:"resource_id,unique"`
	Type       string `json:"service" bun:"service"`
	Name       string `json:"name" bun:"name"`
	Link       []Link `json:"relations" bun:"relations"`
	Provider   string `json:"provider" bun:"provider"`
}
