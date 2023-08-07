package pointers

import (
	"fmt"
	"testing"
)

// A pointer that is assigned to nil is called a nil pointer.
// Compiler assigns nil value to pointer variable at the time of compilation if you do not assign address to pointer variable.
func TestNilPointer(t *testing.T) {
	type Temp struct {
	}
	var pnt *Temp              // pointer
	var intf interface{} = pnt // interface declaration
	//intf = pnt                 // inf is a non-nil interface holding a nil pointer (pnt)

	fmt.Printf("pnt is a nil pointer: %v\n", pnt == nil)
	fmt.Printf("intf is a nil interface: %v\n", intf == nil)
	fmt.Printf("intf is an interface holding a nil pointer: %v\n", intf == (*Temp)(nil))
}
