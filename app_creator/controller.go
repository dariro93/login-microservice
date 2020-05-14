package app_creator

import "login-micro/dataprovider"

type Controller interface {
	StartServer(database dataprovider.Database)
}
