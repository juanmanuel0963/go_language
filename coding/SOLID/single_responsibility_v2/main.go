package main

import (
	"encoding/json"
	"fmt"
	"log"
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

// ----------shape----------
type shape interface {
	area() float64
	name() string
}

// ----------circle----------
type circle struct {
	radius float64
}

// circle implements name
func (c circle) name() string {
	return "circle"
}

// circle implements area
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// ----------square----------
type square struct {
	length float64
}

// square implements name
func (s square) name() string {
	return "square"
}

// square implements area
func (s square) area() float64 {
	return s.length * s.length
}

// ----------outputter----------
type outputter struct {
}

// outputter adds Text function
func (o outputter) Text(s shape) string {
	return fmt.Sprintf("area  of the %s: %f", s.name(), s.area())
}

// outputter adds JSON function
func (o outputter) JSON(s shape) string {
	res := struct {
		Name string  `json:"shape"`
		Area float64 `json:"area"`
	}{
		Name: s.name(),
		Area: s.area(),
	}

	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

// ----------main----------
func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	out := outputter{}
	fmt.Println(out.Text(c))
	fmt.Println(out.JSON(c))
	fmt.Println(out.Text(s))
	fmt.Println(out.JSON(s))
}
