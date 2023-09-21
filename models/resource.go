package models

import "time"

type Resource struct {
	Id         int64             `json:"id" bun:"id,pk,autoincrement"`
	ResourceId string            `json:"resourceId" bun:"resource_id,unique"`
	Provider   string            `json:"provider"`
	Account    string            `json:"account"`
	AccountId  string            `json:"accountId" bun:"account_id"`
	Service    string            `json:"service" bun:"service"`
	Region     string            `json:"region"`
	Name       string            `json:"name" bun:"name"`
	CreatedAt  time.Time         `json:"createdAt" bun:"created_at"`
	FetchedAt  time.Time         `json:"fetchedAt" bun:"fetched_at"`
	Cost       float64           `json:"cost"`
	Metadata   map[string]string `json:"metadata"`
	Relations  []Link            `json:"relations" bun:"relations,default:'[]'"`
	Tags       []Tag             `json:"tags" bun:"tags,default:'[]'"`
	Link       string            `json:"link" bson:"link"`
	Value      string            `bun:",scanonly"` //to be deprecated
}

type Tag struct {
	Key   string `json:"key" bun:"key"`
	Value string `json:"value" bun:"value"`
}

type BulkUpdateTag struct {
	Tags      []Tag `json:"tags"`
	Resources []int `json:"resources"`
}
