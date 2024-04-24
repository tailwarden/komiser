package models

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

func HandleQuery(ctx context.Context, db *bun.DB, queryTitle string, schema interface{}, additionals map[string]string) (sql.Result, error) {
	var resp sql.Result
	var err error
	switch Queries[queryTitle].Type {
	case RAW:
		err = executeRaw(ctx, db, Queries[queryTitle].Query, schema, additionals)
		if err != nil {
			return resp, err
		}
	case SELECT:
		err = executeSelect(ctx, db, Queries[queryTitle].Query, schema, additionals)
		if err != nil {
			return resp, err
		}
	case INSERT:
		resp, err = executeInsert(ctx, db, schema, additionals)
		if err != nil {
			return resp, err
		}
	case DELETE:
		resp, err = executeDelete(ctx, db, schema, Queries[queryTitle].Query, additionals)
		if err != nil {
			return resp, err
		}
	case UPDATE:
		resp, err = executeUpdate(ctx, db, schema, Queries[queryTitle].Query, Queries[queryTitle].Params, additionals)
		if err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func executeRaw(ctx context.Context, db *bun.DB, query string, schema interface{}, additionals map[string]string) error {
	if len(additionals) > 0 {
		query = fmt.Sprintf("%s where", query)
	}

	for key, value := range additionals {
		query = fmt.Sprintf("%s %s = '%s' and", query, key, value)
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

func executeSelect(ctx context.Context, db *bun.DB, query string, schema interface{}, additionals map[string]string) error {
	updatedQuery := db.NewSelect().Model(schema)
	for key, value := range additionals {
		updatedQuery = updatedQuery.Where(fmt.Sprintf("%s = ?", key), value)
	}

	err := updatedQuery.Scan(ctx, schema)
	if err != nil {
		return err
	}
	return nil
}

func executeInsert(ctx context.Context, db *bun.DB, schema interface{}, additionals map[string]string) (sql.Result, error) {
	resp, err := db.NewInsert().Model(schema).Exec(ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func executeDelete(ctx context.Context, db *bun.DB, schema interface{}, query string, additionals map[string]string) (sql.Result, error) {
	resp, err := db.NewDelete().Model(schema).Where("id = ?", additionals["id"]).Exec(ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func executeUpdate(ctx context.Context, db *bun.DB, schema interface{}, query string, columns []string, additionals map[string]string) (sql.Result, error) {
	updatedQuery := db.NewUpdate().Model(schema).Column(columns...)

	for key, value := range additionals {
		updatedQuery = updatedQuery.Where(fmt.Sprintf("%s = ?", key), value)
	}

	updatedQuery = updatedQuery.Returning("*")
	resp, err := updatedQuery.Exec(ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
