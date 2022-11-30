package interfaces

import (
	"fmt"
	"math"
)

// Go interfaces are those that have a defined set of method signatures
// It is a custom type who can take values that has these methods implementation.
// The interfaces are abstract which is why we cannot create its instance.
// But we can create a variable of type interface and that variable can then be
// assigned to a concrete value that has methods required by the interface.
// Due to these reasons, an interface can act as two things:
//- Collection of method signatures
//- Custom types

//Interface also promotes abstraction. In Golang, we can use interfaces for creating
//common abstractions which can be used by multiple types by defining method
//declarations that are compatible with the interface

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func Measure(g geometry) (float64, float64) {

	fmt.Println("Measure")
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

	return g.area(), g.perim()
}
