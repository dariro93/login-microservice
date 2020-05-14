package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"login-micro/controllers/rest/model"
	"login-micro/usecase"
	"github.com/gorilla/mux"
	"strconv"
)

const contentType = "Content-Type"
const applicationJson = "application/json"

var userManager usecase.UsersUseCase

func InjectUsersManager(usrMgr usecase.UsersUseCase) {
	userManager = usrMgr
}

// CreateUserHandler Function handle users creation
func CreateUserHandler(writer http.ResponseWriter, request *http.Request) {
	var userDTO model.UserRequest
	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(reqBody, &userDTO)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(userDTO.String())
	id, err := userManager.CreateUser(userDTO.ToUserEntity())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.Header().Add(contentType, applicationJson)
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("%d", id)))
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	users, err := userManager.GetAllUsers()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.Header().Add(contentType, applicationJson)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(users)
}

func GetUserByID(writer http.ResponseWriter, request *http.Request) {
	idStr := mux.Vars(request)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	userDTO, err := userManager.GetUserByID(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.Header().Add(contentType, applicationJson)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(userDTO)
}
