package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	. "github.com/tailwarden/komiser/models"
	"github.com/uptrace/bun/dialect"
)

func (handler *ApiHandler) NewViewHandler(w http.ResponseWriter, r *http.Request) {
	var view View

	err := json.NewDecoder(r.Body).Decode(&view)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := handler.db.NewInsert().Model(&view).Exec(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	viewId, _ := result.LastInsertId()
	view.Id = viewId

	respondWithJSON(w, 200, view)
}

func (handler *ApiHandler) ListViewsHandler(w http.ResponseWriter, r *http.Request) {
	views := make([]View, 0)

	handler.db.NewRaw("SELECT * FROM views").Scan(handler.ctx, &views)

	respondWithJSON(w, 200, views)
}

func (handler *ApiHandler) GetViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewId, _ := vars["id"]

	view := new(View)
	err := handler.db.NewSelect().Model(view).Where("id = ?", viewId).Scan(handler.ctx)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, 200, view)
}

func (handler *ApiHandler) UpdateViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewId, _ := vars["id"]

	var view View
	err := json.NewDecoder(r.Body).Decode(&view)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = handler.db.NewUpdate().Model(&view).Column("name", "filters", "exclude").Where("id = ?", viewId).Exec(handler.ctx)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Error while updating view")
		return
	}

	respondWithJSON(w, 200, view)
}

func (handler *ApiHandler) DeleteViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewId, _ := vars["id"]

	view := new(View)
	_, err := handler.db.NewDelete().Model(view).Where("id = ?", viewId).Exec(handler.ctx)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error while updating view")
		return
	}

	respondWithJSON(w, 200, "View has been deleted")
}

func (handler *ApiHandler) HideResourcesFromViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewId, _ := vars["id"]

	var view View
	err := json.NewDecoder(r.Body).Decode(&view)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = handler.db.NewUpdate().Model(&view).Column("exclude").Where("id = ?", viewId).Exec(handler.ctx)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Error while updating view")
		return
	}

	respondWithJSON(w, 200, "Resources has been hidden")
}

func (handler *ApiHandler) UnhideResourcesFromViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewId, _ := vars["id"]

	var view View
	err := json.NewDecoder(r.Body).Decode(&view)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = handler.db.NewUpdate().Model(&view).Column("exclude").Where("id = ?", viewId).Exec(handler.ctx)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Error while updating view")
		return
	}

	respondWithJSON(w, 200, "Resources has been revealed")
}

func (handler *ApiHandler) ListHiddenResourcesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewId, _ := vars["id"]

	view := new(View)
	err := handler.db.NewSelect().Model(view).Where("id = ?", viewId).Scan(handler.ctx)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resources := make([]Resource, len(view.Exclude))

	if len(view.Exclude) > 0 {
		s, _ := json.Marshal(view.Exclude)
		handler.db.NewRaw(fmt.Sprintf("SELECT * FROM resources WHERE id IN (%s)", strings.Trim(string(s), "[]"))).Scan(handler.ctx, &resources)

	}

	respondWithJSON(w, 200, resources)
}

func (handler *ApiHandler) GetViewResourcesHandler(w http.ResponseWriter, r *http.Request) {
	var filters []Filter

	limitRaw := r.URL.Query().Get("limit")
	skipRaw := r.URL.Query().Get("skip")
	query := r.URL.Query().Get("query")

	vars := mux.Vars(r)
	viewId, _ := vars["id"]

	view := new(View)
	err := handler.db.NewSelect().Model(view).Where("id = ?", viewId).Scan(handler.ctx)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var limit int64
	var skip int64
	limit = 0
	skip = 0
	l, err := strconv.ParseInt(limitRaw, 10, 64)
	if err != nil {
		limit = 0
	} else {
		limit = l
	}

	s, err := strconv.ParseInt(skipRaw, 10, 64)
	if err != nil {
		skip = 0
	} else {
		skip = s
	}

	err = json.NewDecoder(r.Body).Decode(&filters)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	filterWithTags := false
	whereQueries := make([]string, 0)
	for _, filter := range filters {
		if filter.Field == "name" || filter.Field == "region" || filter.Field == "service" || filter.Field == "provider" || filter.Field == "account" {
			switch filter.Operator {
			case "IS":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("(%s IN (%s))", filter.Field, strings.Join(filter.Values, ","))
				whereQueries = append(whereQueries, query)
			case "IS_NOT":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("(%s NOT IN (%s))", filter.Field, strings.Join(filter.Values, ","))
				whereQueries = append(whereQueries, query)
			case "CONTAINS":
				queries := make([]string, 0)
				specialChar := "%"
				for i := 0; i < len(filter.Values); i++ {
					queries = append(queries, fmt.Sprintf("(%s LIKE '%s%s%s')", filter.Field, specialChar, filter.Values[i], specialChar))
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(%s)", strings.Join(queries, " OR ")))
			case "NOT_CONTAINS":
				queries := make([]string, 0)
				specialChar := "%"
				for i := 0; i < len(filter.Values); i++ {
					queries = append(queries, fmt.Sprintf("(%s NOT LIKE '%s%s%s')", filter.Field, specialChar, filter.Values[i], specialChar))
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(%s)", strings.Join(queries, " AND ")))
			case "IS_EMPTY":
				whereQueries = append(whereQueries, fmt.Sprintf("((coalesce(%s, '') = ''))", filter.Field))
			case "IS_NOT_EMPTY":
				whereQueries = append(whereQueries, fmt.Sprintf("((coalesce(%s, '') != ''))", filter.Field))
			default:
				respondWithError(w, http.StatusBadRequest, "Operation is invalid or not supported")
				return
			}
		} else if strings.HasPrefix(filter.Field, "tag:") {
			filterWithTags = true
			key := strings.ReplaceAll(filter.Field, "tag:", "")
			switch filter.Operator {
			case "CONTAINS":
			case "IS":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' IN (%s)))", key, strings.Join(filter.Values, ","))
				if handler.db.Dialect().Name() == dialect.SQLite {
					query = fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') IN (%s)))", key, strings.Join(filter.Values, ","))
				}
				whereQueries = append(whereQueries, query)
			case "NOT_CONTAINS":
			case "IS_NOT":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' NOT IN (%s)))", key, strings.Join(filter.Values, ","))
				if handler.db.Dialect().Name() == dialect.SQLite {
					query = fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') NOT IN (%s)))", key, strings.Join(filter.Values, ","))
				}
				whereQueries = append(whereQueries, query)
			case "IS_EMPTY":
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') = ''))", key))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' = ''))", key))
				}
			case "IS_NOT_EMPTY":
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') != ''))", key))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' != ''))", key))
				}
			default:
				respondWithError(w, http.StatusBadRequest, "Operation is invalid or not supported")
				return
			}
		} else if filter.Field == "tags" {
			switch filter.Operator {
			case "IS_EMPTY":
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, "json_array_length(tags) = 0")
				} else {
					whereQueries = append(whereQueries, "jsonb_array_length(tags) = 0")
				}
			case "IS_NOT_EMPTY":
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, "json_array_length(tags) != 0")
				} else {
					whereQueries = append(whereQueries, "jsonb_array_length(tags) != 0")
				}
			default:
				respondWithError(w, http.StatusBadRequest, "Operation is invalid or not supported")
				return
			}
		} else if filter.Field == "cost" {
			switch filter.Operator {
			case "EQUAL":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, "The value should be a number")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost = %f)", cost))
			case "BETWEEN":
				min, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, "The value should be a number")
				}
				max, err := strconv.ParseFloat(filter.Values[1], 64)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, "The value should be a number")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost >= %f AND cost <= %f)", min, max))
			case "GREATER_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, "The value should be a number")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost > %f)", cost))
			case "LESS_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, "The value should be a number")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost < %f)", cost))
			default:
				respondWithError(w, http.StatusBadRequest, "Operation is invalid or not supported")
				return
			}
		} else {
			respondWithError(w, http.StatusBadRequest, "Field is invalid or not supported")
			return
		}
	}

	if len(query) > 0 {
		clause := fmt.Sprintf("(name ilike '%s' OR region ilike '%s' OR service ilike '%s' OR provider ilike '%s' OR account ilike '%s' OR tags @> '[{\"value\":\"%s\"}]' or tags @> '[{\"key\":\"%s\"}]')", query, query, query, query, query, query, query)
		whereQueries = append(whereQueries, clause)
	}

	whereClause := strings.Join(whereQueries, " AND ")

	resources := make([]Resource, 0)
	if filterWithTags {
		query := fmt.Sprintf("SELECT id, resource_id, provider, account, service, region, name, created_at, fetched_at,cost, metadata, tags,link FROM resources CROSS JOIN jsonb_array_elements(tags) AS res WHERE %s ORDER BY id LIMIT %d OFFSET %d", whereClause, limit, skip)
		if len(view.Exclude) > 0 {
			s, _ := json.Marshal(view.Exclude)
			query = fmt.Sprintf("SELECT id, resource_id, provider, account, service, region, name, created_at, fetched_at,cost, metadata, tags,link FROM resources CROSS JOIN jsonb_array_elements(tags) AS res WHERE %s AND id NOT IN (%s) ORDER BY id LIMIT %d OFFSET %d", whereClause, strings.Trim(string(s), "[]"), limit, skip)
		}
		if handler.db.Dialect().Name() == dialect.SQLite {
			query = fmt.Sprintf("SELECT resources.id, resources.resource_id, resources.provider, resources.account, resources.service, resources.region, resources.name, resources.created_at, resources.fetched_at, resources.cost, resources.metadata, resources.tags, resources.link FROM resources CROSS JOIN json_each(tags) WHERE type='object' AND %s ORDER BY resources.id LIMIT %d OFFSET %d", whereClause, limit, skip)
			if len(view.Exclude) > 0 {
				s, _ := json.Marshal(view.Exclude)
				query = fmt.Sprintf("SELECT resources.id, resources.resource_id, resources.provider, resources.account, resources.service, resources.region, resources.name, resources.created_at, resources.fetched_at, resources.cost, resources.metadata, resources.tags, resources.link FROM resources CROSS JOIN json_each(tags) WHERE resources.id NOT IN (%s) AND type='object' AND %s ORDER BY resources.id LIMIT %d OFFSET %d", strings.Trim(string(s), "[]"), whereClause, limit, skip)
			}
		}

		fmt.Println(query)
		handler.db.NewRaw(query).Scan(handler.ctx, &resources)
	} else {
		query := fmt.Sprintf("SELECT * FROM resources WHERE %s ORDER BY id LIMIT %d OFFSET %d", whereClause, limit, skip)
		if len(view.Exclude) > 0 {
			s, _ := json.Marshal(view.Exclude)
			query = fmt.Sprintf("SELECT * FROM resources WHERE %s AND id NOT IN (%s) ORDER BY id LIMIT %d OFFSET %d", whereClause, strings.Trim(string(s), "[]"), limit, skip)
		}

		err = handler.db.NewRaw(query).Scan(handler.ctx, &resources)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	respondWithJSON(w, 200, resources)
}
