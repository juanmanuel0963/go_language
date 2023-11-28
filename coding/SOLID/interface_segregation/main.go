package main

import (
	"fmt"
	"math"
)

/*
•	Coming to the “Interface Segregation Principle” …
the goal is to identify the collaborators of a class…
•	The idea is to find new collaborator classes…
and segregate or distribute responsibilities between several classes.
•	For each collaborator, we ask, ¿what methods and attributes does this collaborator require?
*/

// ----------shape----------
type shape interface {
	area() float64
	volume() float64
}

// ----------square----------
type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

// square is a flat surface -> does not have volume
func (s square) volume() float64 {
	return 0
}

// ----------cube----------
type cube struct {
	length float64
}

func (c cube) area() float64 {
	return math.Pow(c.length, 2)
}

func (c cube) volume() float64 {
	return math.Pow(c.length, 3)
}

// ----------areaSum----------
func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

// ----------areaVolumeSum----------
func areaVolumeSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

// ----------main----------
func main() {
	s := square{length: 3}
	c := cube{length: 4}
	fmt.Println(areaSum(s, c))
	fmt.Println(areaVolumeSum(s, c))
}
