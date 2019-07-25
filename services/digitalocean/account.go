package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

type Account struct {
	Email string `json:"email"`
}

func (dg DigitalOcean) DescribeAccount(client *godo.Client) (Account, error) {
	account := Account{}

	a, _, err := client.Account.Get(context.TODO())
	if err != nil {
		return account, err
	}

	account.Email = a.Email

	return account, nil
}
