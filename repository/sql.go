package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
)

type QueryType string

const (
	RAW    QueryType = "RAW"
	SELECT QueryType = "SELECT"
	INSERT QueryType = "INSERT"
	DELETE QueryType = "DELETE"
	UPDATE QueryType = "UPDATE"
)

type Object struct {
	Query  string    `json:"query"`
	Type   QueryType `json:"type"`
	Params []string  `json:"params"`
}

const (
	ListKey                  = "LIST"
	InsertKey                = "INSERT"
	DeleteKey                = "DELETE"
	UpdateAccountKey         = "UPDATE_ACCOUNT"
	UpdateAlertKey           = "UPDATE_ALERT"
	UpdateViewKey            = "UPDATE_VIEW"
	UpdateViewExcludeKey     = "UPDATE_VIEW_EXCLUDE"
	ReScanAccountKey         = "RE_SCAN_ACCOUNT"
	ResourceCountKey         = "RESOURCE_COUNT"
	ResourceCostSumKey       = "RESOURCE_COST_SUM"
	AccountsResourceCountKey = "ACCOUNTS_RESOURCE_COUNT"
	RegionResourceCountKey   = "REGION_RESOURCE_COUNT"
	FilterResourceCountKey   = "FILTER_RESOURCE_COUNT"
	LocationBreakdownStatKey = "LOCATION_BREAKDOWN_STAT"
	UpdateTagsKey            = "UPDATE_TAGS"
	ListRegionsKey           = "LISST_REGIONS"
	ListProvidersKey         = "LIST_PROVIDERS"
	ListServicesKey          = "LIST_SERVICES"
	ListAccountsKey          = "LIST_ACCOUNTS"
)

var Queries = map[string]Object{
	ListKey: {
		Type: SELECT,
	},
	InsertKey: {
		Type: INSERT,
	},
	DeleteKey: {
		Type: DELETE,
	},
	UpdateAccountKey: {
		Type:   UPDATE,
		Params: []string{"name", "provider", "credentials"},
	},
	UpdateAlertKey: {
		Type:   UPDATE,
		Params: []string{"name", "type", "budget", "usage", "endpoint", "secret"},
	},
	UpdateViewKey: {
		Type:   UPDATE,
		Params: []string{"name", "filters", "exclude"},
	},
	UpdateViewExcludeKey: {
		Type:   UPDATE,
		Params: []string{"exclude"},
	},
	ReScanAccountKey: {
		Type:   UPDATE,
		Params: []string{"status"},
	},
	ResourceCountKey: {
		Query: "SELECT COUNT(*) as total FROM resources",
		Type:  RAW,
	},
	ResourceCostSumKey: {
		Query: "SELECT SUM(cost) as sum FROM resources",
		Type:  RAW,
	},
	AccountsResourceCountKey: {
		Query: "SELECT COUNT(*) as count FROM (SELECT DISTINCT account FROM resources) AS temp",
		Type:  RAW,
	},
	RegionResourceCountKey: {
		Query: "SELECT COUNT(*) as count FROM (SELECT DISTINCT region FROM resources) AS temp",
		Type:  RAW,
	},
	FilterResourceCountKey: {
		Query: "SELECT filters as label, COUNT(*) as total FROM resources",
		Type:  RAW,
	},
	LocationBreakdownStatKey: {
		Query: "SELECT region as label, COUNT(*) as total FROM resources GROUP BY region ORDER by total desc;",
		Type:  RAW,
	},
	UpdateTagsKey: {
		Type:   UPDATE,
		Params: []string{"tags"},
	},
	ListRegionsKey: {
		Type:  RAW,
		Query: "SELECT DISTINCT(region) FROM resources",
	},
	ListProvidersKey: {
		Type:  RAW,
		Query: "SELECT DISTINCT(provider) FROM resources",
	},
	ListServicesKey: {
		Type:  RAW,
		Query: "SELECT DISTINCT(service) FROM resources",
	},
	ListAccountsKey: {
		Type:  RAW,
		Query: "SELECT DISTINCT(account) FROM resources",
	},
}

func HandleQuery(ctx context.Context, db *bun.DB, queryTitle string, schema interface{}, conditions [][3]string) (sql.Result, error) {
	var resp sql.Result
	var err error
	query := Queries[queryTitle]
	switch query.Type {
	case RAW:
		err = executeRaw(ctx, db, query.Query, schema, conditions)

	case SELECT:
		err = executeSelect(ctx, db, schema, conditions)

	case INSERT:
		resp, err = executeInsert(ctx, db, schema)

	case DELETE:
		resp, err = executeDelete(ctx, db, schema, conditions)

	case UPDATE:
		resp, err = executeUpdate(ctx, db, schema, query.Params, conditions)
	}
	return resp, err
}

func executeRaw(ctx context.Context, db *bun.DB, query string, schema interface{}, additionals [][3]string) error {
	if len(additionals) > 0 {
		query = fmt.Sprintf("%s where", query)
	}

	for _, triplet := range additionals {
		key, op, value := triplet[0], triplet[1], triplet[2]
		query = fmt.Sprintf("%s %s %s '%s' and", query, key, op, value)
	}

	if len(additionals) > 0 {
		query = query[:len(query)-4]
	}

	err := db.NewRaw(query).Scan(ctx, schema)
	if err != nil {
		return err
	}
	return nil
}

func executeSelect(ctx context.Context, db *bun.DB, schema interface{}, conditions [][3]string) error {
	q := db.NewSelect().Model(schema)

	q = addWhereClause(q.QueryBuilder(), conditions).Unwrap().(*bun.SelectQuery)

	return q.Scan(ctx, schema)
}

func executeInsert(ctx context.Context, db *bun.DB, schema interface{}) (sql.Result, error) {
	resp, err := db.NewInsert().Model(schema).Exec(ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func executeDelete(ctx context.Context, db *bun.DB, schema interface{}, conditions [][3]string) (sql.Result, error) {
	q := db.NewDelete().Model(schema)

	q = addWhereClause(q.QueryBuilder(), conditions).Unwrap().(*bun.DeleteQuery)

	resp, err := q.Exec(ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func executeUpdate(ctx context.Context, db *bun.DB, schema interface{}, columns []string, conditions [][3]string) (sql.Result, error) {
	q := db.NewUpdate().Model(schema).Column(columns...)

	q = addWhereClause(q.QueryBuilder(), conditions).Unwrap().(*bun.UpdateQuery)

	q = q.Returning("*")

	resp, err := q.Exec(ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func addWhereClause(query bun.QueryBuilder, conditions [][3]string) bun.QueryBuilder {
	for _, triplet := range conditions {
		key, op, value := triplet[0], triplet[1], triplet[2]
		query = query.Where(fmt.Sprintf("%s %s ?", key, op), value)
	}
	return query
}
