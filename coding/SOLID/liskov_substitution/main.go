package main

import "fmt"

/*
•	Coming to the “Liskov Substitution Principle” …
the goal is to have superclasses (interfaces) …
able to be replace by subclasses (structs)
without causing problems in the application.
•	It’s the same concept of Polymorphism.
•	The subclass (structs) shares the same signature as its superclass (interfaces).
•	The signature is the collection of all exposed methods and attributes.
*/

// ----------person----------
type person interface {
	getName() string
}

// ----------human----------
type human struct {
	name string
}

func (h human) getName() string {
	return h.name
}

// ----------teacher----------
type teacher struct {
	human
	degree string
	salary float64
}

// ----------student----------
type student struct {
	human
	grades map[string]int
}

// ----------printer----------
type printer struct {
}

func (printer) info(p person) {
	fmt.Println("Name: ", p.getName())
}

// ----------main----------
func main() {
	h := human{name: "Alex"}
	s := student{
		human: human{name: "Mike"},
		grades: map[string]int{
			"Math":    8,
			"English": 9,
		},
	}
	t := teacher{
		human:  human{name: "John"},
		degree: "CS",
		salary: 2000,
	}

	p := printer{}
	p.info(h)
	p.info(s)
	p.info(t)
}
