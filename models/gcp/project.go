package gcp

type Project struct {
	Name       string `json:"name"`
	ID         string `json:"id"`
	CreateTime string `json:"createTime"`
	Number     int64  `json:"number"`
}
