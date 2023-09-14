package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func main() {

	var s Shape = nil

	s = Circle{Radius: 5.0}

	area := s.Area()
	perimeter := s.Perimeter()

	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perimeter)
}
