package main

import (
	"fmt"
)

func safeValues(vals []int, index int) int {
	// deferred recover wrapped with anonymous func
	defer func (){
		if err := recover(); err != nil {
			fmt.Printf("Recovered error: %+v", err)
		}
	}()

	return vals[index]
}

func main() {
	/* panic from runtime case */
	vals := []int {1,23,3,4}
	//v := vals[10]
	//fmt.Println(v)

	/* panic produced from our own code */
	//file, err := os.Open("no-such-file")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//fmt.Println("file open")

	v := safeValues(vals, 10)
	fmt.Println(v)
}
