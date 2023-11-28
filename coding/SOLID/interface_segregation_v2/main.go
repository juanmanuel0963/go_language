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
}

// ----------object----------
type object interface {
	shape
	volume() float64
}

// ----------square----------
type square struct {
	length float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

// ----------cube----------
type cube struct {
	square
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
func areaVolumeSum(shapes ...object) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

// ----------main----------
func main() {
	s1 := square{length: 5}
	s2 := square{length: 6}
	c1 := cube{square: square{length: 3}}
	c2 := cube{square: square{length: 4}}
	fmt.Println(areaSum(s1, s2, c1, c2))
	fmt.Println(areaVolumeSum(c1, c2))
}
