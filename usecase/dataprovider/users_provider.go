package dataprovider

import "login-micro/entities"

// UsersProvider is the interface interface for abstracting different db implementations
// Currently supported ones are postgres and mongo
type UsersProvider interface {
	SaveUser(entityUser entities.User) (int64, error)
	DeleteUser(user entities.User) error
	EditUser(user entities.User) error
	GetUserByID(id int) (*entities.User, error)
	GetAllUsers() (*[]entities.User, error)
	GetUserByName(name string) *entities.User
}
