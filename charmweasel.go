package charmweasel

import (
	"fmt"
)

func Main() int {
	return 0
}

type CustomType int

const (
	Foo CustomType = iota
	Bar
	Baz
)

func VariadicFunction(args ...CustomType) {
	for _, arg := range args {
		switch arg {
		case Foo:
			fmt.Println("Foo")
		case Bar:
			fmt.Println("Bar")
		case Baz:
			fmt.Println("Baz")
		default:
			fmt.Println("Unknown")
		}
	}
}
