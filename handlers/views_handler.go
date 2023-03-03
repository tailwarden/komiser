package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tailwarden/komiser/models"
)

func (handler *ApiHandler) NewViewHandler(w http.ResponseWriter, r *http.Request) {
	var view models.View

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
	views := make([]models.View, 0)

	handler.db.NewRaw("SELECT * FROM views").Scan(handler.ctx, &views)

	respondWithJSON(w, 200, views)
}

func (handler *ApiHandler) UpdateViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewId, _ := vars["id"]

	var view models.View
	err := json.NewDecoder(r.Body).Decode(&view)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = handler.db.NewUpdate().Model(&view).Column("name", "filters", "exclude").Where("id = ?", viewId).Exec(handler.ctx)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error while updating view")
		return
	}

	respondWithJSON(w, 200, view)
}

func (handler *ApiHandler) DeleteViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewId, _ := vars["id"]

	view := new(models.View)
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

	var view models.View
	err := json.NewDecoder(r.Body).Decode(&view)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = handler.db.NewUpdate().Model(&view).Column("exclude").Where("id = ?", viewId).Exec(handler.ctx)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error while updating view")
		return
	}

	respondWithJSON(w, 200, "Resources has been hidden")
}

func (handler *ApiHandler) UnhideResourcesFromViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewId, _ := vars["id"]

	var view models.View
	err := json.NewDecoder(r.Body).Decode(&view)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = handler.db.NewUpdate().Model(&view).Column("exclude").Where("id = ?", viewId).Exec(handler.ctx)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error while updating view")
		return
	}

	respondWithJSON(w, 200, "Resources has been revealed")
}

func (handler *ApiHandler) ListHiddenResourcesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewId, _ := vars["id"]

	view := new(models.View)
	err := handler.db.NewSelect().Model(view).Where("id = ?", viewId).Scan(handler.ctx)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resources := make([]models.Resource, len(view.Exclude))

	if len(view.Exclude) > 0 {
		s, _ := json.Marshal(view.Exclude)
		handler.db.NewRaw(fmt.Sprintf("SELECT * FROM resources WHERE id IN (%s)", strings.Trim(string(s), "[]"))).Scan(handler.ctx, &resources)

	}

	respondWithJSON(w, 200, resources)
}

func (handler *ApiHandler) ListViewAlertsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewId := vars["id"]

	alerts := make([]models.Alert, 0)

	handler.db.NewRaw(fmt.Sprintf("SELECT * FROM alerts WHERE view_id = %s", viewId)).Scan(handler.ctx, &alerts)

	respondWithJSON(w, 200, alerts)
}
