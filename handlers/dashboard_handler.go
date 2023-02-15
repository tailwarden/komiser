package handlers

import (
	"net/http"
)

func (handler *ApiHandler) DashboardStatsHandler(w http.ResponseWriter, r *http.Request) {
	regions := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) as count FROM (SELECT DISTINCT region FROM resources) AS temp").Scan(handler.ctx, &regions)

	resources := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) as count FROM resources").Scan(handler.ctx, &resources)

	cost := struct {
		Sum float64 `bun:"sum" json:"total"`
	}{}

	handler.db.NewRaw("SELECT SUM(cost) as sum FROM resources").Scan(handler.ctx, &cost)

	accounts := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) as count FROM (SELECT DISTINCT account FROM resources) AS temp").Scan(handler.ctx, &accounts)

	output := struct {
		Resources int     `json:"resources"`
		Regions   int     `json:"regions"`
		Costs     float64 `json:"costs"`
		Accounts  int     `json:"accounts"`
	}{
		Resources: resources.Count,
		Regions:   regions.Count,
		Costs:     cost.Sum,
		Accounts:  accounts.Count,
	}

	respondWithJSON(w, 200, output)
}
