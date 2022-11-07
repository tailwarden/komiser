package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	. "github.com/mlabouardy/komiser/models"
	"github.com/uptrace/bun"
)

type ResourcesHandler struct {
	db  *bun.DB
	ctx context.Context
}

func NewResourcesHandler(ctx context.Context, db *bun.DB) *ResourcesHandler {
	handler := ResourcesHandler{
		db:  db,
		ctx: ctx,
	}
	return &handler
}

func (handler *ResourcesHandler) ListResourcesHandler(w http.ResponseWriter, r *http.Request) {
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
		whereClause := fmt.Sprintf("name ilike '%s' OR region ilike '%s' OR service ilike '%s' OR provider ilike '%s' OR account ilike '%s' OR tags @> '[{\"value\":\"%s\"}]' or tags @> '[{\"key\":\"%s\"}]'", query, query, query, query, query, query, query)
		handler.db.NewRaw(fmt.Sprintf("SELECT * FROM resources WHERE %s ORDER BY id LIMIT %d OFFSET %d", whereClause, limit, skip)).Scan(handler.ctx, &resources)
	} else {
		handler.db.NewRaw(fmt.Sprintf("SELECT * FROM resources ORDER BY id LIMIT %d OFFSET %d", limit, skip)).Scan(handler.ctx, &resources)
	}

	respondWithJSON(w, 200, resources)
}

func (handler *ResourcesHandler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	regions := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) FROM (SELECT DISTINCT region FROM resources) AS temp").Scan(handler.ctx, &regions)

	resources := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) FROM resources").Scan(handler.ctx, &resources)

	cost := struct {
		Sum int `bun:"sum" json:"total"`
	}{}

	handler.db.NewRaw("SELECT SUM(count) FROM resources").Scan(handler.ctx, &cost)

	output := struct {
		Resources int `json:"resources"`
		Regions   int `json:"regions"`
		Costs     int `json:"costs"`
	}{
		Resources: resources.Count,
		Regions:   regions.Count,
		Costs:     cost.Sum,
	}

	respondWithJSON(w, 200, output)
}

func (handler *ResourcesHandler) RegionsCounterHandler(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) FROM (SELECT DISTINCT region FROM resources) AS temp").Scan(handler.ctx, &output)

	respondWithJSON(w, 200, output)
}

func (handler *ResourcesHandler) ResourcesCounterHandler(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) FROM resources").Scan(handler.ctx, &output)

	respondWithJSON(w, 200, output)
}

func (handler *ResourcesHandler) CostCounterHandler(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Sum int `bun:"sum" json:"total"`
	}{}

	handler.db.NewRaw("SELECT SUM(count) FROM resources").Scan(handler.ctx, &output)

	respondWithJSON(w, 200, output)
}

func (handler *ResourcesHandler) UpdateTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags := make([]Tag, 0)

	vars := mux.Vars(r)
	resourceId, ok := vars["id"]

	if !ok {
		respondWithError(w, http.StatusBadRequest, "Resource id is missing")
		return
	}

	id, err := strconv.Atoi(resourceId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Resource id should be an integer")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&tags)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resource := Resource{Tags: tags}

	_, err = handler.db.NewUpdate().Model(&resource).Column("tags").Where("id = ?", id).Exec(handler.ctx)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error while updating tags")
		return
	}

	respondWithJSON(w, 200, "Tags has been successfuly updated")
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
