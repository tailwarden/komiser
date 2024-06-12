package controller

import (
	"context"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/repository"
)

func (ctrl *Controller) InsertAlert(c context.Context, alert models.Alert) (alertId int64, err error) {
	alertId, err = ctrl.repo.HandleQuery(c, repository.InsertKey, &alert, nil, "")
	return
}

func (ctrl *Controller) UpdateAlert(c context.Context, alert models.Alert, alertId string) (err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.UpdateAlertKey, &alert, [][3]string{{"id", "=", alertId}}, "")
	return
}

func (ctrl *Controller) DeleteAlert(c context.Context, alertId string) (err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.DeleteKey, new(models.Alert), [][3]string{{"id", "=", alertId}}, "")
	return
}
