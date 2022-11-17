package handlers

import (
	"net/http"
)

func (handler *ApiHandler) StatsHandler(w http.ResponseWriter, r *http.Request) {
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

func (handler *ApiHandler) ListRegionsHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Region string `bun:"region" json:"region"`
	}

	outputs := make([]Output, 0)

	handler.db.NewRaw("SELECT DISTINCT(region) FROM resources").Scan(handler.ctx, &outputs)

	regions := make([]string, 0)

	for _, o := range outputs {
		regions = append(regions, o.Region)
	}

	respondWithJSON(w, 200, regions)
}

func (handler *ApiHandler) ListProvidersHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Provider string `bun:"provider" json:"provider"`
	}

	outputs := make([]Output, 0)

	handler.db.NewRaw("SELECT DISTINCT(provider) FROM resources").Scan(handler.ctx, &outputs)

	providers := make([]string, 0)

	for _, o := range outputs {
		providers = append(providers, o.Provider)
	}

	respondWithJSON(w, 200, providers)
}

func (handler *ApiHandler) ListServicesHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Service string `bun:"service" json:"service"`
	}

	outputs := make([]Output, 0)

	handler.db.NewRaw("SELECT DISTINCT(service) FROM resources").Scan(handler.ctx, &outputs)

	services := make([]string, 0)

	for _, o := range outputs {
		services = append(services, o.Service)
	}

	respondWithJSON(w, 200, services)
}

func (handler *ApiHandler) ListAccountsHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Account string `bun:"account" json:"account"`
	}

	outputs := make([]Output, 0)

	handler.db.NewRaw("SELECT DISTINCT(account) FROM resources").Scan(handler.ctx, &outputs)

	accounts := make([]string, 0)

	for _, o := range outputs {
		accounts = append(accounts, o.Account)
	}

	respondWithJSON(w, 200, accounts)
}

func (handler *ApiHandler) ResourcesCounterHandler(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) FROM resources").Scan(handler.ctx, &output)

	respondWithJSON(w, 200, output)
}

func (handler *ApiHandler) CostCounterHandler(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Sum int `bun:"sum" json:"total"`
	}{}

	handler.db.NewRaw("SELECT SUM(count) FROM resources").Scan(handler.ctx, &output)

	respondWithJSON(w, 200, output)
}

func (handler *ApiHandler) EnableTrackingHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, map[string]bool{"tracking": !handler.noTracking})
}
