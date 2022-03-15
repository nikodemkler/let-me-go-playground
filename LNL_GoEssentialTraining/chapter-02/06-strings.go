package main

import (
	"fmt"
)

func main () {
	book := "That is the book"

	fmt.Printf("book[0] = %v, (type of book[0]: %T)\n\n", book[0], book[0])
	fmt.Printf("book[4] = %v, (type of book[4]: %T)\n\n", book[4], book[4])

	// can't go.. it's immutable!
	//book[5] = 32
	fmt.Printf("Text len: %v\n-------------\n\n", len(book))

	// here I go again on python's slices (ALMOST)
	fmt.Println(book[0:4])
	fmt.Println(book[2:5])
	fmt.Println(book[2:])
	fmt.Println(book[:5])

	// inline concatenation
	fmt.Println(book[1:5] + "!!!")

	// unicode
	unicode := "©"
	fmt.Println(unicode)
	fmt.Println(unicode[0])

	longer := `Hmm
and what 0
	about 1
		indentation 2
			heh? 3
`
	fmt.Println(longer)

	fmt.Println("-------------")
	fmt.Printf(longer)
}