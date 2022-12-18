package functions

import (
	"fmt"
	"testing"
)

// Function closures is a function value that references variables from outside its body.
// The function may access and assign values to the referenced variables.
// For example: adder() returns a closure, which is each bound to its own referenced sum variable.
func TestAdder(t *testing.T) {

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
