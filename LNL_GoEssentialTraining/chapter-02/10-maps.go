package main

import (
	"fmt"
)

func main () {
	stocks := map[string]float64 {
		"foo": 23.4,
		"bar": 99.5,
		"baz": 342.5,
	}

	fmt.Println(len(stocks))
	fmt.Println(stocks)
	fmt.Println("-----------")

	fmt.Printf("\n\non existing key: %v\n", stocks["adasda"])

	// mutable structure
	stocks["qux"] = 2345

	// mutable existing values
	stocks["baz"] = 0

	fmt.Println(len(stocks))
	fmt.Println(stocks)
	fmt.Println("-----------")

	value, keyExists := stocks["qux"]
	fmt.Printf("keyExists=%v, val=%v\n", keyExists, value)
	fmt.Println("-----------")

	value, keyExists = stocks["baz"]
	fmt.Printf("keyExists=%v, val=%v\n", keyExists, value)
	fmt.Println("-----------")

	value, keyExists = stocks["non-existing-keyExists"]
	fmt.Printf("keyExists=%v, val=%v\n", keyExists, value)
	fmt.Println("-----------")

	// delete a key
	delete(stocks, "foo")
	fmt.Println(stocks)

	fmt.Println("-----------")
	for value := range stocks {
		fmt.Println(value)
	}

	fmt.Println("-----------")
	for key, value := range stocks {
		fmt.Printf("%s = %.2f\n", key, value)
	}
}