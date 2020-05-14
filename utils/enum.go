package utils

import "fmt"

type Enum interface {
	fmt.Stringer
	Int() int
}

type EnumContainer interface {
	New() EnumContainer
	Values() []Enum
}

func ToString(enum EnumContainer) string {
	result := "[ "
	for _, e := range enum.Values() {
		result += e.String() + " "
	}
	return result + "]"
}
