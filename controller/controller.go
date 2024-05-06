package controller

import (
	"context"
	"database/sql"

	"github.com/tailwarden/komiser/repository"
)

type Repository interface {
	HandleQuery(context.Context, repository.QueryType, interface{}, [][3]string) (sql.Result, error)
}

type Controller struct {
	repo Repository
}

func New(repo Repository) *Controller {
	return &Controller{repo}
}
