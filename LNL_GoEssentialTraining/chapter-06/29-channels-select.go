package main

import "fmt"

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		ch1 <- 324
	}()
	go func() {
		ch2 <- 333
	}()

	select {
	case val := <- ch1:
		fmt.Printf("Channel1 : %d\n", val)
	case val := <- ch2:
		fmt.Printf("Channel2 : %d\n", val)
	}
}