package controller

import (
	"context"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/repository"
)

func (ctrl *Controller) ListAccounts(c context.Context) (accounts []models.Account, err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.ListKey, &accounts, nil)
	return
}

func (ctrl *Controller) CountResources(c context.Context, provider, name string) (output totalOutput, err error) {
	conditions := [][3]string{}
	if provider != "" {
		conditions = append(conditions, [3]string{"provider", "=", provider})
	}
	if name != "" {
		conditions = append(conditions, [3]string{"account", "=", name})
	}
	_, err = ctrl.repo.HandleQuery(c, repository.ResourceCountKey, &output, conditions)
	return
}

func (ctrl *Controller) InsertAccount(c context.Context, account models.Account) (lastId int64, err error) {
	result, err := ctrl.repo.HandleQuery(c, repository.InsertKey, &account, nil)
	if err != nil {
		return
	}
	return result.LastInsertId()
}

func (ctrl *Controller) RescanAccount(c context.Context, account *models.Account, accountId string) (rows int64, err error) {
	res, err := ctrl.repo.HandleQuery(c, repository.ReScanAccountKey, account, [][3]string{{"id", "=", accountId}, {"status", "=", "CONNECTED"}})
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (ctrl *Controller) DeleteAccount(c context.Context, accountId string) (err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.DeleteKey, new(models.Account), [][3]string{{"id", "=", accountId}})
	return
}

func (ctrl *Controller) UpdateAccount(c context.Context, account models.Account, accountId string) (err error) {
	_, err = ctrl.repo.HandleQuery(c, repository.UpdateAccountKey, &account, [][3]string{{"id", "=", accountId}})
	return
}
