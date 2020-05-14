package controllers

import (
	"login-micro/dataprovider"
	"login-micro/usecase"
	"login-micro/controllers/rest"
	"login-micro/controllers/graphql"
)

func CreateRestController(db dataprovider.Database) error {
	usrProvider := dataprovider.NewUsersProvider(db)
	usrMgr := usecase.NewUsersUseCase(usrProvider)
	rest.InjectUsersManager(usrMgr)
	return rest.ExposeApis()
}

func CreateGraphqlController(db dataprovider.Database) error {
	return graphql.NewGraphqlApp(db).StartApp()
}