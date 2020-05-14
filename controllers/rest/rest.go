package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

func ExposeApis() error {
	router := mux.NewRouter().StrictSlash(true)
	exposeUsersApi(router)
	return http.ListenAndServe(":8081", router)
}

func exposeUsersApi(router *mux.Router) {
	router.HandleFunc("/users", CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/users", GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", GetUserByID).Methods(http.MethodGet)
}