package main

import (
	"fmt"
	"math"
)

type SquareTwo struct {
	Length int
}

func (sq *SquareTwo) Area() float64 {
	return float64(sq.Length * sq.Length)
}

type Circle struct {
	Radius float64
}

func (crc *Circle) Area() float64 {
	return math.Pi * math.Pow(crc.Radius, 2)
}

type Shapes interface {
	Area() float64
}

func sumAreas(shapes []Shapes) float64 {
	var sumOfAll float64
	for _, shape := range shapes {
		sumOfAll += shape.Area()
	}

	return sumOfAll
}

func main() {
	square := &SquareTwo{Length: 10}
	circle := &Circle{Radius: 1}

	var geometrics []Shapes
	geometrics = append(geometrics, square)
	geometrics = append(geometrics, circle)

	fmt.Println("----------")
	fmt.Println(sumAreas(geometrics))
}