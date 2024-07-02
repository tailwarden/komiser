package sqlite

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/repository"
	"github.com/uptrace/bun"
)

type Repository struct {
	mu      sync.RWMutex
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
		Query: "SELECT COUNT(*) as total FROM (SELECT DISTINCT region FROM resources)",
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
	repository.ListResourceWithFilter: {
		Type:  repository.RAW,
		Query: "",
		Params: []string{
			"(name LIKE '%%%s%%' OR region LIKE '%%%s%%' OR service LIKE '%%%s%%' OR provider LIKE '%%%s%%' OR account LIKE '%%%s%%' OR (tags LIKE '%%%s%%'))",
			"SELECT * FROM resources WHERE %s ORDER BY id LIMIT %d OFFSET %d",
			"SELECT * FROM resources ORDER BY id LIMIT %d OFFSET %d",
			"SELECT DISTINCT resources.id, resources.resource_id, resources.provider, resources.account, resources.service, resources.region, resources.name, resources.created_at, resources.fetched_at, resources.cost, resources.metadata, resources.tags, resources.link FROM resources CROSS JOIN json_each(tags) WHERE ",
			"SELECT * FROM resources WHERE %s ORDER BY id LIMIT %d OFFSET %d",
			"SELECT * FROM resources WHERE %s AND id NOT IN (%s) ORDER BY id LIMIT %d OFFSET %d",
		},
	},
	repository.ListRelationWithFilter: {
		Type:  repository.RAW,
		Query: "",
		Params: []string{
			"SELECT DISTINCT resources.resource_id, resources.provider, resources.name, resources.service, resources.relations FROM resources WHERE (json_array_length(relations) > 0)",
		},
	},
	repository.ListStatsWithFilter: {
		Type:  repository.RAW,
		Query: "",
		Params: []string{
			"SELECT COUNT(*) as total FROM (SELECT DISTINCT region FROM resources CROSS JOIN json_each(tags) WHERE type='object' AND %s) AS temp",
			"SELECT COUNT(*) as total FROM resources CROSS JOIN json_each(tags) WHERE type='object' AND %s",
			"SELECT SUM(cost) as sum FROM resources CROSS JOIN json_each(tags) WHERE type='object' AND %s",
			"SELECT COUNT(*) as total FROM (SELECT DISTINCT region FROM resources WHERE %s) AS temp",
			"SELECT COUNT(*) as total FROM resources WHERE %s",
			"SELECT SUM(cost) as sum FROM resources WHERE %s",
		},
	},
}

func (repo *Repository) HandleQuery(ctx context.Context, queryTitle string, schema interface{}, conditions [][3]string, rawQuery string) (resp int64, err error) {
	repo.mu.RLock()
	query, ok := Queries[queryTitle]
	repo.mu.RUnlock()
	if !ok {
		return 0, repository.ErrQueryNotFound
	}
	switch query.Type {
	case repository.RAW:
		if rawQuery != "" && query.Query == "" {
			err = repository.ExecuteRaw(ctx, repo.db, rawQuery, schema, conditions)
		} else {
			err = repository.ExecuteRaw(ctx, repo.db, query.Query, schema, conditions)
		}

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

func (repo *Repository) GenerateFilterQuery(view models.View, queryTitle string, arguments []int64, queryParameter string) ([]string, error) {
	whereQueries := make([]string, 0)
	filterWithTags := false
	for _, filter := range view.Filters {
		switch filter.Field {
		case "account", "resource", "service", "provider", "name", "region":
			query, err := generateStandardFilterQuery(filter, false)
			if err != nil {
				return nil, err
			}
			whereQueries = append(whereQueries, query)
		case "cost":
			query, err := generateCostFilterQuery(filter)
			if err != nil {
				return nil, err
			}
			whereQueries = append(whereQueries, query)
		case "relation":
			query, err := generateRelationFilterQuery(filter)
			if err != nil {
				return nil, err
			}
			whereQueries = append(whereQueries, query)
		case "tags":
			query, err := generateEmptyFilterQuery(filter)
			if err != nil {
				return nil, err
			}
			whereQueries = append(whereQueries, query)
		default:
			if strings.HasPrefix(filter.Field, "tag:") {
				filterWithTags = true
				query, err := generateStandardFilterQuery(filter, true)
				if err != nil {
					return nil, err
				}
				whereQueries = append(whereQueries, query)
			} else {
				return nil, fmt.Errorf("unsupported field: %s", filter.Field)
			}
		}
	}

	whereClause := strings.Join(whereQueries, " AND ")
	return queryBuilderWithFilter(view, queryTitle, arguments, queryParameter, filterWithTags, whereClause), nil
}

func queryBuilderWithFilter(view models.View, queryTitle string, arguments []int64, query string, withTags bool, whereClause string) []string {
	searchQuery := []string{}
	var limit, skip int64
	if len(arguments) >= 2 {
		limit, skip = arguments[0], arguments[1]
	}
	if len(view.Filters) == 0 {
		switch queryTitle {
		case repository.ListRelationWithFilter:
			return append(searchQuery, Queries[queryTitle].Params[0])
		case repository.ListResourceWithFilter:
			tempQuery := ""
			if len(query) > 0 {
				whereClause = fmt.Sprintf(Queries[queryTitle].Params[0], query, query, query, query, query, query, query)
				tempQuery = fmt.Sprintf(Queries[queryTitle].Params[1], whereClause, limit, skip)
			} else {
				tempQuery = fmt.Sprintf(Queries[queryTitle].Params[2], limit, skip)
			}
			return append(searchQuery, tempQuery)
		}
	} else if queryTitle == repository.ListRelationWithFilter {
		return append(searchQuery, Queries[queryTitle].Params[0]+" AND "+whereClause)
	}

	if withTags {
		if queryTitle == repository.ListStatsWithFilter {
			for i := 0; i < 3; i++ {
				searchQuery = append(searchQuery, fmt.Sprintf(Queries[queryTitle].Params[i], whereClause))
			}
			return searchQuery
		}
		tempQuery := fmt.Sprintf(Queries[queryTitle].Params[3]+"type='object' AND %s ORDER BY resources.id LIMIT %d OFFSET %d", whereClause, limit, skip)
		if len(view.Exclude) > 0 {
			s, _ := json.Marshal(view.Exclude)
			tempQuery = fmt.Sprintf(Queries[queryTitle].Params[3]+"resources.id NOT IN (%s) AND type='object' AND %s ORDER BY resources.id LIMIT %d OFFSET %d", whereClause, strings.Trim(string(s), "[]"), limit, skip)
		}
		return append(searchQuery, tempQuery)
	} else {
		if queryTitle == repository.ListStatsWithFilter {
			for i := 3; i < 6; i++ {
				searchQuery = append(searchQuery, fmt.Sprintf(Queries[queryTitle].Params[i], whereClause))
			}
			return searchQuery
		}
		tempQuery := fmt.Sprintf(Queries[queryTitle].Params[4], whereClause, limit, skip)

		if whereClause == "" {
			tempQuery = fmt.Sprintf(Queries[queryTitle].Params[2], limit, skip)
		}

		if len(view.Exclude) > 0 {
			s, _ := json.Marshal(view.Exclude)
			tempQuery = fmt.Sprintf(Queries[queryTitle].Params[5], whereClause, strings.Trim(string(s), "[]"), limit, skip)
		}

		return append(searchQuery, tempQuery)
	}
}

func generateEmptyFilterQuery(filter models.Filter) (string, error) {
	switch filter.Operator {
	case "IS_EMPTY":
		return "json_array_length(tags) = 0", nil
	case "IS_NOT_EMPTY":
		return "json_array_length(tags) != 0", nil
	}
	return "", fmt.Errorf("unsupported operator: %s", filter.Operator)
}

func generateStandardFilterQuery(filter models.Filter, withTag bool) (string, error) {
	key := strings.ReplaceAll(filter.Field, "tag:", "")
	switch filter.Operator {
	case "IS":
		for i := 0; i < len(filter.Values); i++ {
			filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
		}
		if withTag {
			query := fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') IN (%s)))", key, strings.Join(filter.Values, ","))
			return query, nil
		}
		return fmt.Sprintf("(%s IN (%s))", filter.Field, strings.Join(filter.Values, ",")), nil
	case "IS_NOT":
		for i := 0; i < len(filter.Values); i++ {
			filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
		}
		if withTag {
			query := fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') NOT IN (%s)))", key, strings.Join(filter.Values, ","))
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
			return fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') != ''))", key), nil
		}
		return fmt.Sprintf("((coalesce(%s, '') = ''))", filter.Field), nil
	case "IS_NOT_EMPTY":
		if withTag {
			return fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') != ''))", key), nil
		}
		return fmt.Sprintf("((coalesce(%s, '') != ''))", filter.Field), nil
	case "EXISTS":
		return fmt.Sprintf("((json_extract(value, '$.key') = '%s'))", key), nil
	case "NOT_EXISTS":
		return fmt.Sprintf(`(NOT EXISTS (SELECT 1 FROM json_each(resources.tags) WHERE (json_extract(value, '$.key') = '%s')))`, key), nil
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

func generateRelationFilterQuery(filter models.Filter) (string, error) {
	switch filter.Operator {
	case "EQUAL":
		relations, err := strconv.Atoi(filter.Values[0])
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("json_array_length(resources.relations) = %d", relations), err
	case "GREATER_THAN":
		relations, err := strconv.Atoi(filter.Values[0])
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("json_array_length(resources.relations) > %d", relations), err
	case "LESS_THAN":
		relations, err := strconv.Atoi(filter.Values[0])
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("json_array_length(resources.relations) < %d", relations), err
	default:
		return "", fmt.Errorf("unsupported operator: %s", filter.Operator)
	}
}
