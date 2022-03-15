package main

import (
	"fmt"
)

func main () {
	loons := []string {"foo", "bar", "baz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	fmt.Println(len(loons))

	fmt.Println("-------")
	fmt.Println(loons[1])

	fmt.Println("-------")
	fmt.Println(loons[1:])

	fmt.Println("-------")
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	fmt.Println("-------")
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("-------")
	for i, name := range loons {
		fmt.Printf("\niter: %v, val: %v\n", i, name)
	}

	fmt.Println("-------")
	for _, name := range loons {
		fmt.Println(name)
	}

	loons = append(loons, "qux")
	fmt.Println(loons)

	slice := append([]byte("hello "), "world"...)
	fmt.Println(slice)

}