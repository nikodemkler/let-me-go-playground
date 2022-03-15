package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is bigger then 5")
	}

	if x > 100 {
		fmt.Println("X is greater than 100")
	} else {
		fmt.Println("X is lower than 100")
	}

	if x > 5 && x < 15 {
		fmt.Println("X is just OK")
	}

	if x < 20 || x > 30 {
		fmt.Println("X is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}
}
