package app_creator

import (
	"login-micro/dataprovider"
	"login-micro/dataprovider/configuration"
	"login-micro/controllers"
)

type CloudApp struct {
	DBImpl dataprovider.Database
}

func CreateCloudApp() App {
	return &CloudApp{}
}

func (ca *CloudApp) Build() App {
	oracleConnInfo := configuration.NewDBconfiguration("oracle.env")
	ca.DBImpl = dataprovider.NewOracleConn(oracleConnInfo)
	return ca
}

func (ca *CloudApp) StartServer() error {
	return controllers.CreateGraphqlController(ca.DBImpl)
}
