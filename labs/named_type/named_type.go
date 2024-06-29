package named_type

import (
	"fmt"
)

// declaring a new type
type names []string

// declaring a method (function receiver)
func (n names) Print() {
	// n is called method's receiver
	// n is the actual copy of the names we're working with and is available in the function.
	// n is like this or self from OOP.
	// any variable of type names can call this function on itself like variable_name.Print()

	// iterating and printing all names
	for i, name := range n {
		fmt.Println(i, name)
	}
}
