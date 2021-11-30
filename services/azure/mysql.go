package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	. "github.com/mlabouardy/komiser/models/azure"
)

func getMySQLServersClient(subscriptionID string) mysql.ServersClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	serversClient := mysql.NewServersClient(subscriptionID)
	serversClient.Authorizer = a
	return serversClient
}

func getMySQLDatabasesClient(subscriptionID string) mysql.DatabasesClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	databasesClient := mysql.NewDatabasesClient(subscriptionID)
	databasesClient.Authorizer = a
	return databasesClient
}

func (azure Azure) DescribeMySQLInstances(subscriptionID string) ([]MySQLInstance, error) {
	serversClient := getMySQLServersClient(subscriptionID)
	databasesClient := getMySQLDatabasesClient(subscriptionID)
	ctx := context.Background()
	serverListResult, err := serversClient.List(ctx)
	mySqls := make([]MySQLInstance, 0)
	if err != nil {
		return mySqls, err
	}
	rGroups, err := getGroups(subscriptionID)
	if err != nil {
		return mySqls, err
	}
	for _, rGroup := range rGroups {
		servers := *serverListResult.Value
		for _, server := range servers {
			dbListResult, err := databasesClient.ListByServer(ctx, rGroup, *server.Name)
			if err != nil {
				return mySqls, err
			}
			dbs := *dbListResult.Value
			for _, db := range dbs {
				mySqls = append(mySqls, MySQLInstance{
					Name: *db.Name,
					ID:   *db.ID,
				})
			}

		}
	}
	return mySqls, nil
}
