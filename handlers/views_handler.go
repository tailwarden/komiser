package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/mlabouardy/komiser/models"
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
