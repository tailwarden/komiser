package sqlite

import (
	"context"
	"database/sql"

	"github.com/tailwarden/komiser/repository"
	"github.com/uptrace/bun"
)

type Repository struct {
	db      *bun.DB
	queries map[string]repository.Object
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, queries: Queries}
}

var Queries = map[string]repository.Object{
	repository.ListKey: {
		Type: repository.SELECT,
	},
	repository.InsertKey: {
		Type: repository.INSERT,
	},
	repository.DeleteKey: {
		Type: repository.DELETE,
	},
	repository.UpdateAccountKey: {
		Type:   repository.UPDATE,
		Params: []string{"name", "provider", "credentials"},
	},
	repository.UpdateAlertKey: {
		Type:   repository.UPDATE,
		Params: []string{"name", "type", "budget", "usage", "endpoint", "secret"},
	},
	repository.UpdateViewKey: {
		Type:   repository.UPDATE,
		Params: []string{"name", "filters", "exclude"},
	},
	repository.UpdateViewExcludeKey: {
		Type:   repository.UPDATE,
		Params: []string{"exclude"},
	},
	repository.ReScanAccountKey: {
		Type:   repository.UPDATE,
		Params: []string{"status"},
	},
	repository.ResourceCountKey: {
		Query: "SELECT COUNT(*) as total FROM resources",
		Type:  repository.RAW,
	},
	repository.ResourceCostSumKey: {
		Query: "SELECT SUM(cost) as sum FROM resources",
		Type:  repository.RAW,
	},
	repository.AccountsResourceCountKey: {
		Query: "SELECT COUNT(*) as total FROM (SELECT DISTINCT account FROM resources) AS temp",
		Type:  repository.RAW,
	},
	repository.RegionResourceCountKey: {
		Query: "SELECT COUNT(*) as total FROM (SELECT DISTINCT region FROM resources) AS temp",
		Type:  repository.RAW,
	},
	repository.FilterResourceCountKey: {
		Query: "SELECT filters as label, COUNT(*) as total FROM resources",
		Type:  repository.RAW,
	},
	repository.LocationBreakdownStatKey: {
		Query: "SELECT region as label, COUNT(*) as total FROM resources GROUP BY region ORDER by total desc;",
		Type:  repository.RAW,
	},
	repository.UpdateTagsKey: {
		Type:   repository.UPDATE,
		Params: []string{"tags"},
	},
	repository.ListRegionsKey: {
		Type:  repository.RAW,
		Query: "SELECT DISTINCT(region) FROM resources",
	},
	repository.ListProvidersKey: {
		Type:  repository.RAW,
		Query: "SELECT DISTINCT(provider) FROM resources",
	},
	repository.ListServicesKey: {
		Type:  repository.RAW,
		Query: "SELECT DISTINCT(service) FROM resources",
	},
	repository.ListAccountsKey: {
		Type:  repository.RAW,
		Query: "SELECT DISTINCT(account) FROM resources",
	},
}

func (repo *Repository) HandleQuery(ctx context.Context, queryTitle string, schema interface{}, conditions [][3]string) (sql.Result, error) {
	var resp sql.Result
	var err error
	query, ok := Queries[queryTitle]
	if !ok {
		return nil, repository.ErrQueryNotFound
	}
	switch query.Type {
	case repository.RAW:
		err = repository.ExecuteRaw(ctx, repo.db, query.Query, schema, conditions)

	case repository.SELECT:
		err = repository.ExecuteSelect(ctx, repo.db, schema, conditions)

	case repository.INSERT:
		resp, err = repository.ExecuteInsert(ctx, repo.db, schema)

	case repository.DELETE:
		resp, err = repository.ExecuteDelete(ctx, repo.db, schema, conditions)

	case repository.UPDATE:
		resp, err = repository.ExecuteUpdate(ctx, repo.db, schema, query.Params, conditions)
	}
	return resp, err
}
