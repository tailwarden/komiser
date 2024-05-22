package controller

import (
	"context"
	"fmt"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/repository"
)

func (ctrl *Controller) UpdateTags(c context.Context, tags []models.Tag, resourceId string) (resource models.Resource, err error) {
	resource.Tags = tags
	_, err = ctrl.repo.HandleQuery(c, repository.UpdateTagsKey, &resource, [][3]string{{"id", "=", fmt.Sprint(resourceId)}}, "")
	return
}
