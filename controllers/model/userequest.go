package model

import (
	"fmt"
	"login-micro/entities"
)

// UserRequest is the bridge between controller layer and usecase layer
type UserRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// ToUserEntity casts UserRequest struct to entities' user type
func (userDTO *UserRequest) ToUserEntity() entities.User {
	userEntity := entities.User{
		Id:       userDTO.Id,
		Name:     userDTO.Name,
		Lastname: userDTO.Lastname,
		Email:    userDTO.Email,
		Number:   userDTO.Phone,
		Password: userDTO.Password,
	}
	return userEntity
}

func CreateUserDTOFromEntity(user *entities.User) *UserRequest {
	return &UserRequest{
		Id:       user.Id,
		Name:     user.Name,
		Lastname: user.Lastname,
		Email:    user.Email,
		Phone:    user.Number,
	}
}

func (userDTO *UserRequest) String() string {
	return fmt.Sprintf("{id: %d, name: %s, email: %s}", userDTO.Id, userDTO.Name, userDTO.Email)
}
