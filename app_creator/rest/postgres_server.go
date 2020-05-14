package rest

import (
	"login-micro/dataprovider/configuration"
)

type PostgresRestServer struct {
	DBImpl configuration.DBType
	envFilename	string
}

func NewPostgresRestServer() RestApp {
	return &PostgresRestServer{
		DBImpl: configuration.Postgres{},
		envFilename: "postgres.env",
	}
}

func (p *PostgresRestServer) StartApp() error {
	injectDeps(p.envFilename)
	return exposeUsersApi()
}

func (p *PostgresRestServer) RetrieveDBImplName() configuration.DBType {
	return p.DBImpl
}
