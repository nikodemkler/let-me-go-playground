package main

import (
	"fmt"
)

func main () {
	for i := 1; i <= 20; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Printf("fizz buzz")
		} else if i % 3 == 0 {
			fmt.Printf("fizz ")
		} else if i % 5 ==0 {
			fmt.Printf("buzz")
		} else {
			fmt.Printf("%v", i)
		}
		fmt.Printf("\n")
	}

	fmt.Println("------------")

	for i := 1; i <= 20; i++ {

		switch {
		case i % 3 == 0 && i % 5 == 0:
			fmt.Println("fizz buzz")
		case i % 3 == 0:
			fmt.Println("fizz")
		case i % 5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}