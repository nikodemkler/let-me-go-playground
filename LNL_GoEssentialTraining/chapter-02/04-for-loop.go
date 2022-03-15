package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n----\n\n")

	for i := 0; i < 20; i++ {

		if i % 2 != 0{
			fmt.Printf("Odd num %v\n", i)
		}

		if i > 10 {
			break
		}
	}

	fmt.Println("++++++++")
	for i := 0; i < 20; i++ {
		if i < 10 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("++++++++")
	j := 1
	for j < 100 {
		j += 10
		fmt.Println(j)
	}

	fmt.Println("^^^^^^^^")

	k := 0
	for {

		if k > 4 {
			break
		}
		fmt.Println(k)

		k++
	}
}