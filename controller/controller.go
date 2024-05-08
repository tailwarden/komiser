package controller

import (
	"context"
	"database/sql"

	"github.com/tailwarden/komiser/repository"
)

type totalOutput struct {
	Total int `bun:"total" json:"total"`
}

type costOutput struct {
	Total float64 `bun:"sum" json:"total"`
}

type Repository interface {
	HandleQuery(context.Context, repository.QueryType, interface{}, [][3]string) (sql.Result, error)
}

type Controller struct {
	repo Repository
}

func New(repo Repository) *Controller {
	return &Controller{repo}
}
