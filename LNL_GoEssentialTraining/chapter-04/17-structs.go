package main

import "fmt"

type Trade struct {
	Symbol string
	Volume int
	Price float64
	Buy bool
}

func main() {
	t1 := Trade{"MSFT", 100, 5322.23, true}
	fmt.Println(t1)
	fmt.Println(t1.Symbol)

	t2 := Trade{
		Buy: false,
		Symbol: "ABCD",
		Volume: 320,
		Price: 345.54,
	}
	println(t2.Symbol)
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)
}