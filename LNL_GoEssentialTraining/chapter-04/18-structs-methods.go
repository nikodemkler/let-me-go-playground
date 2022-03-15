package main

import "fmt"

type Trade2 struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *Trade2) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}

type FooPoint struct {
	X int
	Y int
}

func (p *FooPoint) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	t1 := Trade2{"MSFT", 10, 99.98, true}
	fmt.Printf("%+v\n", t1.Value())

	p := &FooPoint{1, 2}
	p.Move(10, 10)

	fmt.Println(p)
}
