package app_creator

import (
	"login-micro/utils"
	"fmt"
)

type AppType utils.Enum

type (
	OnPremAppType struct{}
	CloudAppType struct{}

)

func (OnPremAppType) Int() int {
	return 1
}

func (OnPremAppType) String() string {
	return "ON PREMISE APP"
}

func (CloudAppType) Int() int {
	return 2
}

func (CloudAppType) String() string {
	return "CLOUD APP"
}

func RetrieveAppTypeByName(name string) (AppType, error) {
	switch (name) {
	case "onprem":
		return OnPremAppType{}, nil
	case "cloud":
		return CloudAppType{}, nil
	default:
		return nil, fmt.Errorf("There is no such app type: %s", name)
	}
}
