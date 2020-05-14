package model

import (
	"fmt"
	"login-micro/entities"
)

// User model for database struct
type User struct {
	Id       int `db:"u_id"`
	Name     string `db:"u_name"`
	Lastname string `db:"u_lastname"`
	Email    string `db:"u_email"`
	Password string `db:"u_password"`
	Number   string `db:"u_phone"`
}

func NewFromEntityUser(userEntity entities.User) User {
	return User{
		Id:       userEntity.Id,
		Name:     userEntity.Name,
		Lastname: userEntity.Lastname,
		Email:    userEntity.Email,
		Password: userEntity.Password,
		Number:   userEntity.Number,
	}
}

// Cast DB user model to User entity
func (u User) ToUserEntity() *entities.User {
	return &entities.User{
		Id:       u.Id,
		Name:     u.Name,
		Lastname: u.Lastname,
		Number:   u.Number,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u User) String() string {
	return fmt.Sprintf("{id: %d, name: %s, lastname: %s, email: %s, number: %s}", u.Id, u.Name, u.Lastname, u.Email, u.Number)
}
