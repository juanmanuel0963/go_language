package miscellaneous

import (
	"fmt"
	"testing"
)

// If a variable is not assigned any value, go automatically initialises it with the zero value of the variable's type.
// Go is strongly typed, so variables declared as belonging to one type cannot be assigned a value of another type.
func TestVariableDeclaration(t *testing.T) {

	// 1 - variable declaration, then assignment
	var age int
	age = 29

	fmt.Println(age)

	// 2 - variable declaration with initial value
	var age2 int = 29
	fmt.Println(age2)

	// 3 - Type inference
	var age3 = 29
	fmt.Println(age3)

	// 4 - declaring multiple variables
	var width, height int = 100, 50
	fmt.Println(width)
	fmt.Println(height)
	fmt.Println(100)
	fmt.Println(50)

	//Here, we are assigning values of a type Integer number, Floating-Point number and
	//string to the three variables in a single line of code.
	var a, b, c = 3, 4.5, "foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 5 - declare variables belonging to different types in a single statement
	var (
		name1 = "initialvalue1"
		name2 = "initialvalue2"
	)

	fmt.Println(name1)
	fmt.Println(name2)

	// 6 - short hand declaration

	name, age4 := "naveen", 29 //short hand declaration
	petName := "apolo"

	fmt.Println(name)
	fmt.Println(age4)
	fmt.Println(petName)
}
