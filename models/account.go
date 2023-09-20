package models

type Account struct {
	Id          int64             `json:"id" bun:"id,pk,autoincrement"`
	Provider    string            `json:"provider"`
	Name        string            `json:"name"`
	Credentials map[string]string `json:"credentials" bun:"credentials,unique"`
	Status      string            `json:"status"`
	Resources   int               `json:"resources" bun:",scanonly"`
}
