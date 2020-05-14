package usecase

import (
	"errors"
	"fmt"
	"login-micro/entities"
	"regexp"
)

var emailRegexp *regexp.Regexp = regexp.MustCompile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)")

// ValidateUserParams validate that user is valid
func ValidateUserParams(user entities.User) error {
	fmt.Println(user)
	if user.Name == "" {
		return throwError("Name cannot be empty")
	}
	if user.Lastname == "" {
		return throwError("Lastname cannot be empty")
	}
	if user.Email == "" {
		return throwError("Email cannot be empty")
	}
	if !emailRegexp.Match([]byte(user.Email)) {
		return throwError("Email pattern does not match a valid one")
	}
	if len(user.Password) < 10 {
		return throwError("Password minimun allowed length is 10 characters")
	}
	return nil
}

func throwError(msg string) error {
	return errors.New(msg)
}
