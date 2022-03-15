package main

import "fmt"

func main () {
	nums := []int {16, 8, 42, 4, 23, 15}
	var maxValue int

	// BTW vars "declaration" syntax also initialize it..
	var uintInitEx uint
	var stringInitEx string
	var boolInitEx bool

	fmt.Println(uintInitEx)
	fmt.Println(stringInitEx)
	fmt.Println(boolInitEx)

	for _, num := range nums {
		if num > maxValue {
			maxValue = num
		}
	}
	fmt.Println(maxValue)


	// alternatively
	fmt.Println("------------")
	maxValue = nums[0]
	for _, num := range nums[1:] {
		if num > maxValue {
			maxValue = num
		}
	}

	fmt.Println(maxValue)
}