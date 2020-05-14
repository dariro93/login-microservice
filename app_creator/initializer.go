package app_creator

import (
	"fmt"
	"github.com/joho/godotenv"
	"login-micro/configuration"
)

type AppConfig struct {
	appType AppType
}

type App interface {
	Build() App
	StartServer() error
}

func CreateApp() App {
	appConfig := loadAppConfiguration()
	switch appConfig.appType.(type) {
	case OnPremAppType:
		return CreateOnPremApp().Build()
	case CloudAppType:
		return CreateCloudApp().Build()
	default:
		panic(fmt.Sprintf("There is not such application type: %s", appConfig.appType))
	}
}

func loadAppConfiguration() AppConfig {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	var appConfig AppConfig
	if err != nil {
		panic(err)
	}
	appType, err := RetrieveAppTypeByName(configuration.RetrieveProperty("app.type"))
	appConfig.appType = appType
	return appConfig
}
