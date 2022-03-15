package main

import (
	"fmt"
	"math"
)

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value %f", num)
	}
	return math.Sqrt(num), nil
}

func main() {
	s1, err1 := sqrt(9)
	if err1 != nil {
		fmt.Printf("ERROR: %s\n", err1)
	} else {
		fmt.Println(s1)
	}

	s2, err2 := sqrt(-1)
	if err2 != nil {
		fmt.Printf("ERROR: %s\n", err2)
	} else {
		fmt.Println(s2)
	}
}