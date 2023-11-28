package main

import (
	"fmt"
	"math"
)

/*
•	Coming to the “Open Closed Principle” …
the goal is to have classes (structs) …
close to modifications…
but open to future extensions.
•	It promotes the idea of creating a subclass…
if I need to add new features to a class…
instead of modifying the current class.
•	If I need to modify a class…
this can introduce a destructive ripple...
that might have unintended consequences.
*/

// ----------shape----------
type shape interface {
	area() float64
}

// ----------circle----------
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// ----------square----------
type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

// ----------triangle----------
type triangle struct {
	height float64
	base   float64
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

// ----------calculator----------
type calculator struct {
}

func (a calculator) areaSum(shapes ...shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}

// ----------main----------
func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	t := triangle{height: 3, base: 7}
	calc := calculator{}
	fmt.Println("area sum:", calc.areaSum(c, s, t))
}
