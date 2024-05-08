package controller

import (
	"context"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/repository"
)

func (ctrl *Controller) LocationStatsBreakdown(c context.Context) (groups []models.OutputResources, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.LocationBreakdownStatKey, &groups, [][3]string{})
	return
}
