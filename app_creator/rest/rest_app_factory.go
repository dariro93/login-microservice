package rest

import (
	"login-micro/dataprovider/configuration"
)

type RestApp interface {
	StartApp() error
	RetrieveDBImplName() configuration.DBType
}

func NewRestApp(dbtype configuration.DBType) RestApp {
	switch dbtype.(type) {
	case configuration.Postgres:
		return NewPostgresRestServer()
	case configuration.Oracle:
		return NewOracleRestServer()
	default:
		panic("There is not such db type supported")
		return nil
	}
}
