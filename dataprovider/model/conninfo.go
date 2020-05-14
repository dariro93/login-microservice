package model

import (
	"fmt"
	"gopai/utils"
)

type ConnectionInfo struct {
	DBName   string
	Port     int
	Password string
	Host     string
	User     string
	DBType   utils.Enum
}

func (c *ConnectionInfo) ToConnectionChain() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d", c.DBName, c.User, c.Password, c.Host, c.Port)
}
