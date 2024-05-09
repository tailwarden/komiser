package controller

import (
	"context"
	"strings"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/repository"
)

func (ctrl *Controller) GetResource(c context.Context, resourceId string) (resource models.Resource, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListKey, &resource, [][3]string{{"resource_id", "=", resourceId}})
	return
}

func (ctrl *Controller) GetResources(c context.Context, idList string) (resources []models.Resource, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListKey, &resources, [][3]string{{"id", "IN", "(" + strings.Trim(idList, "[]") + ")"}})
	return
}

func (ctrl *Controller) ListResources(c context.Context) (resources []models.Resource, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListKey, &resources, [][3]string{})
	return
}

func (ctrl *Controller) CountRegionsFromResources(c context.Context) (regions totalOutput, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.RegionResourceCountKey, &regions, [][3]string{})
	return
}

func (ctrl *Controller) CountRegionsFromAccounts(c context.Context) (accounts totalOutput, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.AccountsResourceCountKey, &accounts, [][3]string{})
	return
}

func (ctrl *Controller) SumResourceCost(c context.Context) (cost costOutput, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ResourceCostSumKey, &cost, [][3]string{})
	return
}
