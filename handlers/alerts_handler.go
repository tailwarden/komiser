package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tailwarden/komiser/models"
)

func (handler *ApiHandler) IsSlackEnabledHandler(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Enabled bool `json:"enabled"`
	}{
		Enabled: false,
	}
	if len(handler.cfg.Slack.Webhook) > 0 {
		output.Enabled = true
	}

	respondWithJSON(w, 200, output)
}

func (handler *ApiHandler) NewAlertHandler(w http.ResponseWriter, r *http.Request) {
	var alert models.Alert

	err := json.NewDecoder(r.Body).Decode(&alert)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := handler.db.NewInsert().Model(&alert).Exec(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	alertId, _ := result.LastInsertId()
	alert.Id = alertId

	respondWithJSON(w, 200, alert)
}

func (handler *ApiHandler) UpdateAlertHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alertId := vars["id"]

	var alert models.Alert
	err := json.NewDecoder(r.Body).Decode(&alert)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = handler.db.NewUpdate().Model(&alert).Column("name", "type", "budget", "usage").Where("id = ?", alertId).Exec(handler.ctx)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error while updating alert")
		return
	}

	respondWithJSON(w, 200, alert)
}

func (handler *ApiHandler) DeleteAlertHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alertId := vars["id"]

	alert := new(models.Alert)
	_, err := handler.db.NewDelete().Model(alert).Where("id = ?", alertId).Exec(handler.ctx)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error while updating alert")
		return
	}

	respondWithJSON(w, 200, "Alert has been deleted")
}
