package main

import (
	"fmt"
)

// slice maps passed by reference
func doubleAt(values []int, index int) {
	values[index] *= 2
}

func double(num int) {
	num *= 2
}

func doublePointer(num *int) {
	*num *= 2
}

func main() {
	numbers := []int {2,4,1,7,5}
	fmt.Println("before")
	fmt.Println(numbers)

	doubleAt(numbers, 1)

	fmt.Println("after")
	fmt.Println(numbers)

	// passing by value (as default for basic types)
	sampleNum := 23
	double(sampleNum)
	fmt.Println(sampleNum)

	// passing by reference
	doublePointer(&sampleNum)
	fmt.Println(sampleNum)
}