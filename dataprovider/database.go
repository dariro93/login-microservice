package dataprovider

import (
	"login-micro/entities"
	dbmodel "login-micro/dataprovider/model"
)

type Database interface {
	CreateConnection(conn *dbmodel.ConnectionInfo) Database
	SaveUser(user entities.User) (int64, error)
	DeleteUser(id int) error
	EditUser(user entities.User) error
	GetUserByID(id int) (*entities.User, error)
	GetUserByName(name string) *entities.User
	GetAllUsers() (*[]dbmodel.User, error)
}
