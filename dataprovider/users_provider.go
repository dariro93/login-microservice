package dataprovider

import (
	"login-micro/entities"
)

type userProviderImpl struct {
	dbConnection Database
}

func NewUsersProvider(dbConnection Database) *userProviderImpl {
	return &userProviderImpl{dbConnection: dbConnection}
}

func (usrepo userProviderImpl) DeleteUser(user entities.User) error {
	return usrepo.dbConnection.DeleteUser(user.Id)
}

func (usrepo userProviderImpl) EditUser(user entities.User) error {
	return usrepo.dbConnection.EditUser(user)
}

func (usrepo userProviderImpl) GetUserByID(id int) (*entities.User, error) {
	return usrepo.dbConnection.GetUserByID(id)
}

func (usrepo userProviderImpl) GetUserByName(name string) *entities.User {
	return usrepo.dbConnection.GetUserByName(name)
}

func (usrepo userProviderImpl) SaveUser(user entities.User) (int64, error) {
	return usrepo.dbConnection.SaveUser(user)
}

func (usrepo userProviderImpl) GetAllUsers() (*[]entities.User,  error) {
	users, err := usrepo.dbConnection.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var entityUsers []entities.User
	for _, user := range *users {
		entityUsers = append(entityUsers, *user.ToUserEntity())
	}
	return &entityUsers, nil
}
