package configuration

import (
	"fmt"
	"login-micro/utils"
	"strings"
)

type DBType utils.Enum

type (
	Postgres struct{}
	Mongo    struct{}
	Oracle 	 struct{}
)

func (Postgres) Int() int {
	return 1
}

func (pg Postgres) String() string {
	return "POSTGRES"
}

func (Mongo) Int() int {
	return 2
}

func (m Mongo) String() string {
	return "MONGO"
}

func (Oracle) Int() int {
	return 3
}

func (o Oracle) String() string {
	return "ORACLE"
}

func RetrieveDBTypeByName(name string) (utils.Enum, error) {
	switch strings.TrimSpace(strings.ToLower(name)) {
	case "postgres":
		return Postgres{}, nil
	case "mongo":
		return Mongo{}, nil
	case "oracle":
		return Oracle{}, nil

	}
	return nil, fmt.Errorf("There is no such database type: %s", name)
}

func Values() []utils.Enum {
	return []utils.Enum{Postgres{}, Mongo{}}
}
