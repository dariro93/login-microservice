package dataprovider

import (
	"fmt"
	dbmodel "login-micro/dataprovider/model"
	"login-micro/dataprovider/configuration"
)

func Build(connInfo *dbmodel.ConnectionInfo) Database {
	switch connInfo.DBType.(type) {
	case configuration.Postgres:
		return NewPostgresConn(connInfo)
	case configuration.Oracle:
		return NewOracleConn(connInfo)
	default:
		panic(fmt.Sprintf("Database %s is not supported", connInfo.DBType))
	}
}
