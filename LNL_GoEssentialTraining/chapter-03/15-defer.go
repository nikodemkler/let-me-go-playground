package main

import (
	"fmt"
)

func cleanUp(name string) {
	fmt.Printf("Clean up: %s\n\n", name)
}

func worker() {
	defer cleanUp("B")
	defer cleanUp("C")
	fmt.Printf("worker\n\n")
}

func main() {
	defer cleanUp("A")

	fmt.Println("before worker")
	worker()
	fmt.Println("after worker")
	defer cleanUp("D")
}
