package handlers

import "net/http"

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

func (handler *ApiHandler) RegionsCounterHandler(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) FROM (SELECT DISTINCT region FROM resources) AS temp").Scan(handler.ctx, &output)

	respondWithJSON(w, 200, output)
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
