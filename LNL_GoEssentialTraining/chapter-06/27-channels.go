// channels
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// This will block
	/*
		<-ch
		fmt.Println("Here")
	*/

	go func() {
		// Send number of the channel
		ch <- 353
	}()

	val := <-ch
	fmt.Printf("Channel val: %d\n", val)

	fmt.Println("---------")

	go func() {
		for i :=0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		 val := <-ch
		 fmt.Printf("received %d\n", val)
	}
	fmt.Println("--------")
	go func() {
		for j :=0; j < 3; j++ {
			fmt.Printf("sending %d\n", j)
			ch <- j
			time.Sleep(time.Second * 2)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("Channel received: %d\n", i)
	}
}