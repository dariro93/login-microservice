package app_creator

import (
	"login-micro/dataprovider"
	"login-micro/dataprovider/configuration"
	"login-micro/controllers"
)

type OnPremApp struct {
	DBImpl dataprovider.Database
}

func CreateOnPremApp() App {
	return &OnPremApp{}
}

func (op *OnPremApp) Build() App {
	postgresConnInfo := configuration.NewDBconfiguration("postgres.env")
	op.DBImpl = dataprovider.NewPostgresConn(postgresConnInfo)
	return op
}

func (op *OnPremApp) StartServer() error {
	return controllers.CreateRestController(op.DBImpl)
}