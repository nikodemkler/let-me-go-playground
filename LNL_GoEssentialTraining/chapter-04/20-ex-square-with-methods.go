package main

import (
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func (pnt *Point) Move(dx int, dy int) {
	pnt.X += dx
	pnt.Y += dy
}

type Square struct {
	Center Point
	Length int
}

func (sqr *Square) Move(dx int, dy int) {
	sqr.Center.Move(dx, dy)
}

func (sqr *Square) Area() int {
	return sqr.Length * sqr.Length
}

func NewSquare(x int, y int, length int) (*Square, error) {

	if length <= 0 {
		return nil, fmt.Errorf("length needs to greater than zero")
	}

	return &Square{
		Center: Point{X:x, Y:y},
		Length: length,
	}, nil
}

func main() {
	newSquare, err := NewSquare(1, 2, 5)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Printf("%+v\n", newSquare)
	fmt.Printf("Area: %+v\n", newSquare.Area())
}
