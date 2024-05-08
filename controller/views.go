package controller

import (
	"context"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/repository"
)

func (ctrl *Controller) GetView(c context.Context, viewId string) (view models.View, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListKey, &view, [][3]string{{"id", "=", viewId}})
	return
}
