package controller

import (
	"context"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/repository"
)

func (ctrl *Controller) LocationStatsBreakdown(c context.Context) (groups []models.OutputResources, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.LocationBreakdownStatKey, &groups, [][3]string{}, "")
	return
}

func (ctrl *Controller) ListRegions(c context.Context) (regions []regionOutput, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListRegionsKey, &regions, nil, "")
	return
}

func (ctrl *Controller) ListProviders(c context.Context) (providers []providerOutput, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListProvidersKey, &providers, nil, "")
	return
}

func (ctrl *Controller) ListServices(c context.Context) (services []serviceOutput, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListServicesKey, &services, nil, "")
	return
}

func (ctrl *Controller) ListAccountNames(c context.Context) (accounts []accountOutput, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListAccountsKey, &accounts, nil, "")
	return
}

func (ctrl *Controller) StatsWithFilter(c context.Context, view models.View, arguments []int64, queryParameter string) (regionCount totalOutput, resourceCount totalOutput, costCount costOutput, err error) {
	queries, err := ctrl.repo.GenerateFilterQuery(view, repository.ListStatsWithFilter, arguments, queryParameter)
	if err != nil {
		return
	}
	_, err = ctrl.repo.HandleQuery(c, repository.ListStatsWithFilter, &regionCount, [][3]string{}, queries[0])
	if err != nil {
		return
	}

	// for resource count
	_, err = ctrl.repo.HandleQuery(c, repository.ListStatsWithFilter, &resourceCount, [][3]string{}, queries[1])
	if err != nil {
		return
	}

	// for cost sum
	_, err = ctrl.repo.HandleQuery(c, repository.ListStatsWithFilter, &costCount, [][3]string{}, queries[2])
	if err != nil {
		return
	}
	return
}