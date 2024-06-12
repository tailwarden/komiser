package controller

import (
	"context"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/repository"
)

func (ctrl *Controller) GetView(c context.Context, viewId string) (view models.View, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListKey, &view, [][3]string{{"id", "=", viewId}}, "")
	return
}

func (ctrl *Controller) ListViews(c context.Context) (views []models.View, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListKey, &views, [][3]string{}, "")
	return
}

func (ctrl *Controller) InsertView(c context.Context, view models.View) (viewId int64, err error) {
	viewId, err = ctrl.repo.HandleQuery(c, repository.InsertKey, &view, nil, "")
	return
}

func (ctrl *Controller) UpdateView(c context.Context, view models.View, viewId string) (err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.UpdateViewKey, &view, [][3]string{{"id", "=", viewId}}, "")
	return
}

func (ctrl *Controller) DeleteView(c context.Context, viewId string) (err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.DeleteKey, new(models.View), [][3]string{{"id", "=", viewId}}, "")
	return
}

func (ctrl *Controller) UpdateViewExclude(c context.Context, view models.View, viewId string) (err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.UpdateViewExcludeKey, &view, [][3]string{{"id", "=", viewId}}, "")
	return
}

func (ctrl *Controller) ListViewAlerts(c context.Context, viewId string) (alerts []models.Alert, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListKey, &alerts, [][3]string{{"view_id", "=", viewId}}, "")
	return
}
