package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	handler.db.NewRaw(fmt.Sprintf("SELECT * FROM resources LIMIT %d OFFSET %d", limit, skip)).Scan(handler.ctx, &resources)

	respondWithJSON(w, 200, resources)
}

func (handler *ResourcesHandler) RegionsCounterHandler(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Count int `bun:"count", json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) FROM (SELECT DISTINCT region FROM resources) AS temp").Scan(handler.ctx, &output)

	respondWithJSON(w, 200, output)
}

func (handler *ResourcesHandler) ResourcesCounterHandler(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Count int `bun:"count", json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) FROM resources").Scan(handler.ctx, &output)

	respondWithJSON(w, 200, output)
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
