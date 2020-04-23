package main

import (
	"fmt"
	"math"
)

type Figure interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	a float64
}
type Circle struct {
	r float64
}

func (s Square) Area() float64 {
	return s.a * s.a
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (s Square) Perimeter() float64 {
	return 4 * s.a
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func main() {
	var s Figure = Square{7}
	var c Figure = Circle{5}

	fmt.Println("The area of a Square:", s.Area())
	fmt.Println("The perimeter of the Square:", s.Perimeter())
	fmt.Println("The area of a Circle:", c.Area())
	fmt.Println("The perimeter of a Circle:", c.Perimeter())
}
