package main

import (
	"fmt"
	"math"
)

/*
•	Coming to the “Single Responsibility Principle” …
the goal is to have a class (struct) …
which can be summarized with a succinct responsibility statement.
•	The class (struct) should have responsibility over a single part of the software…
and that responsibility should be entirely encapsulated by the class (struct).
•	In the context of a class (struct) handling “Events” …
a possible statement could be:
o	The “Events” class (struct) is responsible for creating and modifying events…
ensuring accurate scheduling of personal and corporate events and due dates.
*/

// ----------circle----------
type circle struct {
	radius float64
}

func (c circle) area() {
	fmt.Printf("circle area: %f\n", math.Pi*c.radius*c.radius)
}

// ----------square----------
type square struct {
	length float64
}

func (s square) area() {
	fmt.Printf("square area: %f\n", s.length*s.length)
}

// ----------main----------
func main() {
	c := circle{radius: 5}
	c.area()

	s := square{length: 7}
	s.area()
}
