package graphql

import (
	"fmt"
	"login-micro/dataprovider/configuration"
)

type Graphql struct {

}

func NewGraphqlApp(dbtype configuration.DBType) *Graphql {
	fmt.Println("Not implemented yet...")
	return &Graphql{}
}

func (g *Graphql) StartApp() error {
	fmt.Println("Starting a Graphql server")
	return nil
}
