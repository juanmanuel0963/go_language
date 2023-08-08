package named_type

import "testing"

//Create a Go program that defines a named type and a method (receiver function) for that type.
func TestNamedType(t *testing.T) {

	// declaring a value of type names
	friends := names{"Dan", "Marry"}
	// calling the print() method on friends variable
	friends.Print()

	// another way to call a method on a named type
	names.Print(friends)
}
