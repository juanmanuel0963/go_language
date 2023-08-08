package structs

import (
	"testing"
)

// Empty Struct - When building an object, and only being interested in a grouping of methods and no intermediary data,
// or when you do not plan to retain the object state.
// In the example below it does not make a difference whether the method is called on the same (case #1)
// or on two different objects (case #2):

// Case #1.
func TestStructsFunc_v1(t *testing.T) {

	var lamp Lamp
	lamp.On()
	lamp.Off()
	// Output:
	// on
	// off

}

// Case #2.
func TestStructsFunc_v2(t *testing.T) {

	Lamp{}.On()
	Lamp{}.Off()
	// Output:
	// on
	// off
}

//How do we perform inheritance with Golang?
//------------------------------------------------
//- This is a bit of a trick question: there is no inheritance in Golang because it does not support classes.
//- However, you can mimic inheritance behavior using composition to use an existing struct object to define a starting behavior of a new object.
//- Once the new object is created, functionality can be extended beyond the original struct.

//The Lamp struct contains On(), Off() functions.
//These functions are embedded into the child struct Street by simply listing the struct at the top of the implementation of Street.

func TestStructsInheritance_v1(t *testing.T) {

	var street Street
	street.On()
	street.Off()
	// Output:
	// on
	// off
}
