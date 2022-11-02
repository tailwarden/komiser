package handlers

import (
	"context"
	"encoding/json"
	"net/http"

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

	handler.db.NewRaw("SELECT * FROM resources").Scan(handler.ctx, &resources)

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
