package rest

import (
	"github.com/gorilla/mux"
	"login-micro/controllers"
	"net/http"
	"login-micro/dataprovider"
	"login-micro/dataprovider/configuration"
	"login-micro/usecase"
)

func exposeUsersApi() error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", controllers.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/users", controllers.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controllers.GetUserByID).Methods(http.MethodGet)
	return http.ListenAndServe(":8081", router)
}

func injectDeps(envFilename string) {
	dbConnection := dataprovider.Build(configuration.NewDBconfiguration(envFilename))
	usrProvider := dataprovider.NewUsersProvider(dbConnection)
	usrMgr := usecase.NewUsersUseCase(usrProvider)
	controllers.InjectUsersManager(usrMgr)
}
