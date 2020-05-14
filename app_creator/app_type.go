package app_creator

import (
	"login-micro/utils"
	"fmt"
)

type AppType utils.Enum

type (
	RestAppType struct{}
	GraphqlAppType struct{}
)

func (RestAppType) Int() int {
	return 1
}

func (RestAppType) String() string {
	return "RESTAPP"
}

func (GraphqlAppType) Int() int {
	return 2
}

func (GraphqlAppType) String() string {
	return "GRAPHQL APP"
}

func RetrieveAppTypeByName(name string) (AppType, error) {
	switch (name) {
	case "rest":
		return RestAppType{}, nil
	default:
		return nil, fmt.Errorf("There is no such app type: %s", name)
	}
}
