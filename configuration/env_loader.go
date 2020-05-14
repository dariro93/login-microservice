package configuration

import (
	"os"
	"fmt"
)

func RetrieveProperty(name string) string {
	property := os.Getenv(name)
	if property == "" {
		panic(fmt.Sprintf("Configuration error: please provide a valid property %v value\n", name))
	}
	return property
}
