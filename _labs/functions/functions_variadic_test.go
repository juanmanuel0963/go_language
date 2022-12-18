package functions

import (
	"fmt"
	"testing"
)

// The function that takes a variable number of arguments is called a variadic function.
// We can pass zero or more parameters in the variadic function. The best example of a
// variadic function is fmt.Printf which requires one fixed argument as the first
// parameter and it can accept any arguments.
func TestJoinStringVariadicFunction(t *testing.T) {

	// To demonstrate zero argument
	sJoinString := JoinStringVariadicFunction()
	fmt.Println(sJoinString)

	// To demonstrate multiple arguments
	sJoinString = JoinStringVariadicFunction("Interview", "Bit")
	fmt.Println(sJoinString)

	sJoinString = JoinStringVariadicFunction("Golang", "Interview", "Questions")
	fmt.Println(sJoinString)
}
