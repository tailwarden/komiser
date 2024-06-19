package repository

import (
	"context"
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
	ListResourceWithFilter   = "LIST_RESOURCE_WITH_FILTER"
	ListRelationWithFilter   = "LIST_RELATION_WITH_FILTER"
	ListStatsWithFilter      = "LIST_STATS_WITH_FILTER"
)

func ExecuteRaw(ctx context.Context, db *bun.DB, query string, schema interface{}, additionals [][3]string) error {
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

func ExecuteSelect(ctx context.Context, db *bun.DB, schema interface{}, conditions [][3]string) error {
	q := db.NewSelect().Model(schema)

	q = addWhereClause(q.QueryBuilder(), conditions).Unwrap().(*bun.SelectQuery)

	return q.Scan(ctx, schema)
}

func ExecuteInsert(ctx context.Context, db *bun.DB, schema interface{}) (id int64, err error) {
	res, err := db.NewInsert().Model(schema).Returning("id").Exec(ctx, &id)
	if err != nil {
		_id, err := res.LastInsertId()
		if err != nil {
			id = _id
		}
	}
	return
}

func ExecuteDelete(ctx context.Context, db *bun.DB, schema interface{}, conditions [][3]string) (int64, error) {
	q := db.NewDelete().Model(schema)

	q = addWhereClause(q.QueryBuilder(), conditions).Unwrap().(*bun.DeleteQuery)

	resp, err := q.Exec(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func ExecuteUpdate(ctx context.Context, db *bun.DB, schema interface{}, columns []string, conditions [][3]string) (int64, error) {
	q := db.NewUpdate().Model(schema).Column(columns...)

	q = addWhereClause(q.QueryBuilder(), conditions).Unwrap().(*bun.UpdateQuery)

	q = q.Returning("*")

	resp, err := q.Exec(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func addWhereClause(query bun.QueryBuilder, conditions [][3]string) bun.QueryBuilder {
	for _, triplet := range conditions {
		key, op, value := triplet[0], triplet[1], triplet[2]
		query = query.Where(fmt.Sprintf("%s %s ?", key, op), value)
	}
	return query
}
