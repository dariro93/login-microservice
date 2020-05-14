package rest

import "login-micro/dataprovider/configuration"

type OracleRestServer struct {
	DBImpl configuration.DBType
	envFilename string
}

func NewOracleRestServer() RestApp {
	return &OracleRestServer{
		DBImpl: configuration.Oracle{},
		envFilename: "oracle.env",
	}
}

func (o OracleRestServer) StartApp() error {
	injectDeps(o.envFilename)
	return exposeUsersApi()
}

func (o *OracleRestServer) RetrieveDBImplName() configuration.DBType {
	return o.DBImpl
}