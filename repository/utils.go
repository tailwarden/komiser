package repository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tailwarden/komiser/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect"
)

func generateFilterQuery(db *bun.DB, filters []models.Filter, queryBuilder func([]string) (string, error)) (string, error) {
	whereQueries := make([]string, 0)
	for _, filter := range filters {
		switch filter.Field {
		case "account", "resource", "provider", "name", "region":
			query, err := generateStandardFilterQuery(db, filter, false)
			if err != nil {
				return "", err
			}
			whereQueries = append(whereQueries, query) 
		case "cost":
			query, err := generateCostFilterQuery(filter)
			if err != nil {
				return "", err
			}
			whereQueries = append(whereQueries, query) 
		case "relation":
			query, err := generateRelationFilterQuery(db, filter)
			if err != nil {
				return "", err
			}
			whereQueries = append(whereQueries, query)
		case "tags:":
			query, err := generateStandardFilterQuery(db, filter, true)
			if err != nil {
				return "", err
			}
			whereQueries = append(whereQueries, query)
		case "tags":
			query, err := generateEmptyFilterQuery(db, filter)
			if err != nil {
				return "", err
			}
			whereQueries = append(whereQueries, query)
		default:
			return "", fmt.Errorf("unsupported field: %s", filter.Field)
		}
	}
	return queryBuilder(whereQueries)
}

func generateEmptyFilterQuery(db *bun.DB, filter models.Filter) (string, error) {
	switch filter.Operator {
	case "IS_EMPTY":
		if db.Dialect().Name() == dialect.SQLite {
			return "json_array_length(tags) = 0", nil
		} else {
			return "jsonb_array_length(tags) = 0", nil
		}
	case "IS_NOT_EMPTY":
		if db.Dialect().Name() == dialect.SQLite {
			return "json_array_length(tags) != 0", nil
		} else {
			return "jsonb_array_length(tags) != 0", nil
		}
	}
	return "", fmt.Errorf("unsupported operator: %s", filter.Operator)
}

func generateStandardFilterQuery(db *bun.DB, filter models.Filter, withTag bool) (string, error) {
	for i := 0; i < len(filter.Values); i++ {
		filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
	}
	key := strings.ReplaceAll(filter.Field, "tag:", "")
	switch filter.Operator {
	case "IS":
		if withTag {
			query := fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' IN (%s)))", key, strings.Join(filter.Values, ","))
			if db.Dialect().Name() == dialect.SQLite {
				query = fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') IN (%s)))", key, strings.Join(filter.Values, ","))
			}
			return query, nil
		}
		return fmt.Sprintf("(%s IN (%s))", filter.Field, strings.Join(filter.Values, ",")), nil
	case "IS_NOT":
		if withTag {
			query := fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' NOT IN (%s)))", key, strings.Join(filter.Values, ","))
			if db.Dialect().Name() == dialect.SQLite {
				query = fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') NOT IN (%s)))", key, strings.Join(filter.Values, ","))
			}
			return query, nil
		}
		return fmt.Sprintf("(%s NOT IN (%s))", filter.Field, strings.Join(filter.Values, ",")), nil
	case "CONTAINS":
		queries := make([]string, 0)
		specialChar := "%"
		for i := 0; i < len(filter.Values); i++ {
			queries = append(queries, fmt.Sprintf("(%s LIKE '%s%s%s')", filter.Field, specialChar, filter.Values[i], specialChar))
		}
		return fmt.Sprintf("(%s)", strings.Join(queries, " OR ")), nil
	case "NOT_CONTAINS":
		queries := make([]string, 0)
		specialChar := "%"
		for i := 0; i < len(filter.Values); i++ {
			queries = append(queries, fmt.Sprintf("(%s NOT LIKE '%s%s%s')", filter.Field, specialChar, filter.Values[i], specialChar))
		}
		return fmt.Sprintf("(%s)", strings.Join(queries, " AND ")), nil
	case "IS_EMPTY":
		if withTag {
			if db.Dialect().Name() == dialect.SQLite {
				return fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') != ''))", key), nil
			} else {
				return fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' != ''))", key), nil
			}
		}
		return fmt.Sprintf("((coalesce(%s, '') = ''))", filter.Field), nil
	case "IS_NOT_EMPTY":
		if withTag {
			if db.Dialect().Name() == dialect.SQLite {
				return fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') != ''))", key), nil
			} else {
				return fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' != ''))", key), nil
			}
		}
		return fmt.Sprintf("((coalesce(%s, '') != ''))", filter.Field), nil
	case "EXISTS":
		if db.Dialect().Name() == dialect.SQLite {
			return fmt.Sprintf("((json_extract(value, '$.key') = '%s'))", key), nil
		} else {
			return fmt.Sprintf("((res->>'key' = '%s'))", key), nil
		}
	case "NOT_EXISTS":
		if db.Dialect().Name() == dialect.SQLite {
			return fmt.Sprintf(`(NOT EXISTS (SELECT 1 FROM json_each(resources.tags) WHERE (json_extract(value, '$.key') = '%s')))`, key), nil
		} else {
			return fmt.Sprintf("((res->>'key' != '%s'))", key), nil
		}
	default:
		return "", fmt.Errorf("unsupported operator: %s", filter.Operator)
	}
}

func generateCostFilterQuery(filter models.Filter) (string, error) {
	switch filter.Operator {
	case "EQUAL":
		value, err := strconv.ParseFloat(filter.Values[0], 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("(cost = %f)", value), nil
	case "BETWEEN":
		min, err := strconv.ParseFloat(filter.Values[0], 64)
		if err != nil {
			return "", err
		}
		max, err := strconv.ParseFloat(filter.Values[1], 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("(cost >= %f AND cost <= %f)", min, max), nil		
	case "GREATER_THAN":
		cost, err := strconv.ParseFloat(filter.Values[0], 64)
		if err != nil {
			return "", err
		}		
		return fmt.Sprintf("(cost > %f)", cost), err
	case "LESS_THAN":
		cost, err := strconv.ParseFloat(filter.Values[0], 64)
		if err != nil {
			return "", err
		}		
		return fmt.Sprintf("(cost < %f)", cost), nil
	default:
		return "", fmt.Errorf("unsupported operator for cost field: %s", filter.Operator)
	}
}

func generateRelationFilterQuery(db *bun.DB, filter models.Filter) (string, error) {
	switch filter.Operator {
	case "EQUAL":
		relations, err := strconv.Atoi(filter.Values[0])
		if err != nil {
			return "", err
		}
		if db.Dialect().Name() == dialect.SQLite {
			return fmt.Sprintf("json_array_length(resources.relations) = %d", relations), err
		} else {
			return fmt.Sprintf("jsonb_array_length(resources.relations) = %d", relations), err
		}
	case "GREATER_THAN":
		relations, err := strconv.Atoi(filter.Values[0])
		if err != nil {
			return "", err
		}
		if db.Dialect().Name() == dialect.SQLite {
			return fmt.Sprintf("json_array_length(resources.relations) > %d", relations), err
		} else {
			return fmt.Sprintf("jsonb_array_length(resources.relations) > %d", relations), err
		}
	case "LESS_THAN":
		relations, err := strconv.Atoi(filter.Values[0])
		if err != nil {
			return "", err
		}
		if db.Dialect().Name() == dialect.SQLite {
			return fmt.Sprintf("json_array_length(resources.relations) < %d", relations), err
		} else {
			return fmt.Sprintf("jsonb_array_length(resources.relations) < %d", relations), err
		}
	default:
		return "", fmt.Errorf("unsupported operator: %s", filter.Operator)
	}
}

func AppendResourceQuery(whereQueries []string) (string, error) {
	
	return "", nil
}