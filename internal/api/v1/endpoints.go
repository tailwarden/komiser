package v1

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlabouardy/komiser/handlers"
	"github.com/uptrace/bun"
)

func Endpoints(ctx context.Context, noTracking bool, db *bun.DB) *mux.Router {
	r := mux.NewRouter()

	api := handlers.NewApiHandler(ctx, noTracking, db)

	r.HandleFunc("/resources", api.ListResourcesHandler)
	r.HandleFunc("/resources/search", api.FilterResourcesHandler).Methods("POST")
	r.HandleFunc("/resources/tags", api.BulkUpdateTagsHandler).Methods("POST")
	r.HandleFunc("/resources/count", api.ResourcesCounterHandler)
	r.HandleFunc("/resources/{id}/tags", api.UpdateTagsHandler).Methods("POST")

	r.HandleFunc("/views", api.ListViewsHandler).Methods("GET")
	r.HandleFunc("/views", api.NewViewHandler).Methods("POST")
	r.HandleFunc("/views/{id}", api.GetViewHandler).Methods("GET")
	r.HandleFunc("/views/{id}", api.UpdateViewHandler).Methods("PUT")
	r.HandleFunc("/views/{id}", api.DeleteViewHandler).Methods("DELETE")

	r.HandleFunc("/regions", api.ListRegionsHandler)
	r.HandleFunc("/providers", api.ListProvidersHandler)
	r.HandleFunc("/services", api.ListServicesHandler)
	r.HandleFunc("/accounts", api.ListAccountsHandler)
	r.HandleFunc("/costs", api.CostCounterHandler)
	r.HandleFunc("/stats", api.StatsHandler)
	r.HandleFunc("/stats/search", api.FilterStatsHandler).Methods("POST")
	r.HandleFunc("/tracking", api.EnableTrackingHandler)

	r.PathPrefix("/").Handler(http.FileServer(assetFS()))

	return r
}
