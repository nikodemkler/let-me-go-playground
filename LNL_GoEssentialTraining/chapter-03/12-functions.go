package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a/b, a % b
}

func main() {
	sumValue := add(1,2)
	fmt.Println(sumValue)

	div, modulo := divmod(6,4)
	fmt.Printf("div=%d, modulo=%d\n", div, modulo)

	aaa, _ := divmod(6,4)
	fmt.Println(aaa)
}