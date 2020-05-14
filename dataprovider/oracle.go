package dataprovider

import (
	dbmodel "login-micro/dataprovider/model"
	"login-micro/entities"
	"fmt"
)

type Oracle struct {
	data   map[int]interface{}
	lastID int
}

func NewOracleConn(connInfo *dbmodel.ConnectionInfo) Database {
	var oracle Oracle
	return oracle.CreateConnection(connInfo)
}

func (o *Oracle) CreateConnection(conn *dbmodel.ConnectionInfo) Database {
	o.data = make(map[int]interface{})
	o.lastID = 0
	return o
}

func (o *Oracle) SaveUser(user entities.User) (int64, error) {
	dbUser := dbmodel.NewFromEntityUser(user)
	storedID := o.lastID
	dbUser.Id = o.lastID
	o.data[storedID] = dbUser
	o.lastID++
	return int64(storedID), nil
}

func (o *Oracle) DeleteUser(id int) error {
	delete(o.data, id)
	return nil
}

func (o *Oracle) EditUser(user entities.User) error {
	userI := o.data[user.Id]
	if userI == nil {
		return fmt.Errorf("There is not any user stored with id %d", user.Id)
	}
	userI = dbmodel.NewFromEntityUser(user)
	return nil
}

func (o *Oracle) GetUserByID(id int) (usr *entities.User, err error) {
	defer func() {
		if r := recover(); r != nil {
			usr = nil
			err = fmt.Errorf("There is not any user with ID %d", id)
		}
	}()
	var user dbmodel.User
	user = o.data[id].(dbmodel.User)
	return user.ToUserEntity(), nil
}

func (o *Oracle) GetUserByName(name string) *entities.User {
	return nil
}

func (o *Oracle) GetAllUsers() (*[]dbmodel.User, error) {
	var users []dbmodel.User
	for _, user := range o.data {
		users = append(users, user.(dbmodel.User))
	}
	return &users, nil
}
