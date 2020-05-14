package usecase

import (
	"login-micro/entities"
	"login-micro/usecase/dataprovider"
	controllerModel "login-micro/controllers/rest/model"
)

// UsersUseCase is the interface exporting users' manager behaviour as a dependency
type UsersUseCase interface {
	CreateUser(user entities.User) (int64, error)
	GetUserByID(id int) (*controllerModel.UserRequest, error)
	GetAllUsers() ([]controllerModel.UserRequest, error)
}

// UsersUserCaseImpl is the use case service. Independent from infrastructure services
type UsersUserCaseImpl struct {
	usersProvider dataprovider.UsersProvider
}

func NewUsersUseCase(usersProvider dataprovider.UsersProvider) *UsersUserCaseImpl {
	um := &UsersUserCaseImpl{
		usersProvider: usersProvider,
	}
	return um
}

// CreateUser tells the repository to store current user
func (um *UsersUserCaseImpl) CreateUser(user entities.User) (int64, error) {
	err := ValidateUserParams(user)
	if err != nil {
		return 0, err
	}
	return um.usersProvider.SaveUser(user)
}

func (um *UsersUserCaseImpl) GetUserByID(id int) (*controllerModel.UserRequest, error) {
	user, err := um.usersProvider.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	userDTO := controllerModel.CreateUserDTOFromEntity(user)
	return userDTO, nil
}

func (um *UsersUserCaseImpl) GetAllUsers() ([]controllerModel.UserRequest, error) {
	entityUsers, err := um.usersProvider.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var controllerUsers []controllerModel.UserRequest
	for _, entityUser := range *entityUsers {
		controllerUsers = append(controllerUsers, *controllerModel.CreateUserDTOFromEntity(&entityUser))
	}
	return controllerUsers, nil
}
