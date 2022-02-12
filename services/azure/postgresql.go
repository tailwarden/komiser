package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	. "github.com/mlabouardy/komiser/models/azure"
)

func getPostgreSQLServerClient(subscriptionID string) postgresql.ServersClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	pgsServerClient := postgresql.NewServersClient(subscriptionID)
	pgsServerClient.Authorizer = a
	return pgsServerClient
}

func getPSQLDatabasesClient(subscriptionID string) postgresql.DatabasesClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	databasesClient := postgresql.NewDatabasesClient(subscriptionID)
	databasesClient.Authorizer = a
	return databasesClient
}

func (azure Azure) DescribePostgreSQLInstances(subscriptionID string) ([]PostgreSQL, error) {
	serversClient := getPostgreSQLServerClient(subscriptionID)
	databasesClient := getPSQLDatabasesClient(subscriptionID)
	ctx := context.Background()
	serverListResult, err := serversClient.List(ctx)
	psqls := make([]PostgreSQL, 0)
	if err != nil {
		return psqls, err
	}
	rGroups, err := getGroups(subscriptionID)
	if err != nil {
		return psqls, err
	}
	for _, rGroup := range rGroups {
		servers := *serverListResult.Value
		for _, server := range servers {
			dbListResult, err := databasesClient.ListByServer(ctx, rGroup, *server.Name)
			if err != nil {
				return psqls, err
			}
			dbs := *dbListResult.Value
			for _, db := range dbs {
				psqls = append(psqls, PostgreSQL{
					Name: *db.Name,
					ID:   *db.ID,
				})
			}
		}
	}
	return psqls, nil
}
