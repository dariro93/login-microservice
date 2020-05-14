package app_creator

import (
	"fmt"
	"github.com/joho/godotenv"
	dbConf "login-micro/dataprovider/configuration"
	"login-micro/configuration"
	"login-micro/app_creator/rest"
	"login-micro/app_creator/graphql"
)

type AppConfig struct {
	appType AppType
	dbimpl  dbConf.DBType
}

type App interface {
	StartApp() error
}

func CreateApp() App {
	appConfig := loadAppConfiguration()
	switch appConfig.appType.(type) {
	case RestAppType:
		return rest.NewRestApp(appConfig.dbimpl)
	case GraphqlAppType:
		return graphql.NewGraphqlApp(appConfig.dbimpl)
	default:
		panic(fmt.Sprintf("There is not such application type %s", appConfig.appType))
	}
}

func loadAppConfiguration() AppConfig {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	var appConfig AppConfig
	dbType, err := dbConf.RetrieveDBTypeByName(configuration.RetrieveProperty("db.impl"))
	if err != nil {
		panic(err)
	}
	appType, err := RetrieveAppTypeByName(configuration.RetrieveProperty("app.type"))
	appConfig.appType = appType
	appConfig.dbimpl = dbType
	return appConfig
}
