package mcq

import (
	"fmt"
	"testing"
)

// What is the output of the following code snippet?
func TestMCQ_q1(t *testing.T) {
	x := 3
	y := &x

	fmt.Println(*y)

	*y = 4

	fmt.Println(x)
}

// Anonymous functions
// Let's start with a simple example which assigns a function to a variable.
func TestMCQ_q2(t *testing.T) {

	a := func() {
		fmt.Println("hello world first class function")
	}
	fmt.Println("Let's print function")
	a()
	fmt.Printf("%T", a)

}

// Closure functions
// Closures are a special case of anonymous functions.
// Closures are anonymous functions that access the variables defined outside the body of the function.
func TestMCQ_q3(t *testing.T) {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

// Closure functions
// Closures are a special case of anonymous functions.
// Closures are anonymous functions that access the variables defined outside the body of the function.
func TestMCQ_q4(t *testing.T) {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}

// User defined function types
// The code snippet below creates a new function type add which accepts two integer arguments and returns a integer.
// Now we can define variables of type add.
func TestMCQ_q5(t *testing.T) {

	type add func(a int, b int) int

	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s)

}
func simple1(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

// Higher-order functions
// Passing functions as arguments to other functions
func TestMCQ_q6(t *testing.T) {

	f := func(a, b int) int {
		return a + b
	}
	simple1(f)

}

func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

// Higher-order functions
// Returning functions from other functions
func TestMCQ_q7(t *testing.T) {

	s := simple2()
	fmt.Println(s(60, 7))

}

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

// Practical use of first class functions
// We will create a program that filters a slice of students based on some criteria.
// Let's approach this step by step.
func TestMCQ_q8(t *testing.T) {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

// This program will perform the same operations on each element of a slice and return the result.
// For example, if we want to multiply all integers in a slice by 5 and return the output, it can be easily done using first class functions.
func TestMCQ_q9(t *testing.T) {
	a := []int{5, 6, 7, 8, 9}
	r := iMap(a, func(n int) int {
		return n * 5
	})
	fmt.Println(r)
}
