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

// ----------circle----------
type circle struct {
	radius float64
}

// ----------square----------
type square struct {
	length float64
}

// ----------calculator----------
type calculator struct {
}

func (a calculator) areaSum(shapes ...interface{}) float64 {
	var sum float64
	for _, shape := range shapes {
		switch shape := shape.(type) {
		case circle:
			r := shape.radius
			sum += math.Pi * r * r
		case square:
			l := shape.length
			sum += l * l
		}
	}
	return sum
}

// ----------main----------
func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	calc := calculator{}
	fmt.Println("area sum:", calc.areaSum(c, s))
}
