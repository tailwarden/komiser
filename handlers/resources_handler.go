package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	. "github.com/mlabouardy/komiser/models"
	"github.com/uptrace/bun"
)

type ApiHandler struct {
	db         *bun.DB
	ctx        context.Context
	noTracking bool
}

func NewApiHandler(ctx context.Context, noTracking bool, db *bun.DB) *ApiHandler {
	handler := ApiHandler{
		db:         db,
		ctx:        ctx,
		noTracking: noTracking,
	}
	return &handler
}

func (handler *ApiHandler) ListResourcesHandler(w http.ResponseWriter, r *http.Request) {
	resources := make([]Resource, 0)

	limitRaw := r.URL.Query().Get("limit")
	skipRaw := r.URL.Query().Get("skip")
	query := r.URL.Query().Get("query")

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

	if len(query) > 0 {
		whereClause := fmt.Sprintf("(name ilike '%s' OR region ilike '%s' OR service ilike '%s' OR provider ilike '%s' OR account ilike '%s' OR tags @> '[{\"value\":\"%s\"}]' or tags @> '[{\"key\":\"%s\"}]')", query, query, query, query, query, query, query)
		handler.db.NewRaw(fmt.Sprintf("SELECT * FROM resources WHERE %s ORDER BY id LIMIT %d OFFSET %d", whereClause, limit, skip)).Scan(handler.ctx, &resources)
	} else {
		handler.db.NewRaw(fmt.Sprintf("SELECT * FROM resources ORDER BY id LIMIT %d OFFSET %d", limit, skip)).Scan(handler.ctx, &resources)
	}

	respondWithJSON(w, 200, resources)
}

func (handler *ApiHandler) FilterResourcesHandler(w http.ResponseWriter, r *http.Request) {
	var filters []Filter

	limitRaw := r.URL.Query().Get("limit")
	skipRaw := r.URL.Query().Get("skip")
	query := r.URL.Query().Get("query")

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
				whereQueries = append(whereQueries, query)
			case "NOT_CONTAINS":
			case "IS_NOT":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' NOT IN (%s)))", key, strings.Join(filter.Values, ","))
				whereQueries = append(whereQueries, query)
			case "IS_EMPTY":
				whereQueries = append(whereQueries, fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' = ''))", key))
			case "IS_NOT_EMPTY":
				whereQueries = append(whereQueries, fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' != ''))", key))
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
		handler.db.NewRaw(fmt.Sprintf("SELECT * FROM resources CROSS JOIN jsonb_array_elements(tags) AS res WHERE %s ORDER BY id LIMIT %d OFFSET %d", whereClause, limit, skip)).Scan(handler.ctx, &resources)
	} else {
		err = handler.db.NewRaw(fmt.Sprintf("SELECT * FROM resources WHERE %s ORDER BY id LIMIT %d OFFSET %d", whereClause, limit, skip)).Scan(handler.ctx, &resources)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	respondWithJSON(w, 200, resources)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
