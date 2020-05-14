package configuration

import (
	"strconv"

	"github.com/joho/godotenv"
	dbmodel "login-micro/dataprovider/model"
	"login-micro/configuration"
)

func NewDBconfiguration(envFilename string) *dbmodel.ConnectionInfo {
	err := godotenv.Load(envFilename)
	if err != nil {
		panic(err)
	}
	dbConf := &dbmodel.ConnectionInfo{
		Host:     configuration.RetrieveProperty("db.host"),
		DBName:   configuration.RetrieveProperty("db.name"),
		Password: configuration.RetrieveProperty("db.password"),
		User:     configuration.RetrieveProperty("db.username"),
	}
	dbType, err := RetrieveDBTypeByName(configuration.RetrieveProperty("db.type"))
	if err != nil {
		panic(err)
	}
	dbConf.DBType = dbType
	port, err := strconv.Atoi(configuration.RetrieveProperty("db.port"))
	if err != nil {
		panic(err)
	}
	dbConf.Port = port
	return dbConf
}
