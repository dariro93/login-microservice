package main

import "login-micro/app_creator"

func init() {
	/*container := container.RetrieveContainer()
	dbConnection := dataprovider.Build(configuration.NewDBconfiguration())
	container.RegisterInstance("dbConn", dbConnection)
	usrProvider := dataprovider.NewUsersProvider(dbConnection)
	container.RegisterInstance("usrProvider", usrProvider)
	usrMgr := usecase.NewUserMgr(usrProvider)
	container.RegisterInstance("usrMgr", usrMgr)
	controllers.InjectUsersManager(usrMgr)*/
}

func main() {
	/*container := container.RetrieveContainer()
	var dbconnection interface{}
	dbconnection, _ = container.RetrieveInstanceByType("dbConn")
	fmt.Println(reflect.TypeOf(dbconnection))*/

	/*router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", controllers.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/users", controllers.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controllers.GetUserByID).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8081", router))*/

	/*app := app_creator.CreateApp()
	app.StartApp()*/

	app := app_creator.CreateApp()
	app.StartServer()
}
