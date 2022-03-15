package main

import (
	"fmt"
	"os"
)

type Foo struct {
	Bar int
	Baz string
}

// FooFoo aka constructor equivalent
func FooFoo(baz string, bar int) (*Foo, error) {

	if len(baz) < 3 {
		return nil, fmt.Errorf("baz needs to be min 3 characters long")
	}

	return &Foo{
		Baz: baz,
		Bar: bar,
	}, nil
}

func main() {
	newFoo, err := FooFoo("sds", 123)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(newFoo)
}
