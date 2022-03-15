package main

import (
	"fmt"
)

func main() {
	x := 4
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}

	switch {
	default:
		fmt.Println("Less then 2")
	case x > 100:
		fmt.Println("Greater than 100")
	case x > 2:
		fmt.Println("Greater than 2")
	case x > 1:
		fmt.Println("Greater than 1")
	}
}