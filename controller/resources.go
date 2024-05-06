package controller

import (
	"context"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/repository"
)

func (ctrl *Controller) GetResource(c context.Context, resourceId string) (resource models.Resource, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListKey, &resource, [][3]string{{"resource_id", "=", resourceId}})
	return
}
