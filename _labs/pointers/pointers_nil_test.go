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
	var pnt *Temp       // pointer
	var inf interface{} // interface declaration
	inf = pnt           // inf is a non-nil interface holding a nil pointer (pnt)

	fmt.Printf("pnt is a nil pointer: %v\n", pnt == nil)
	fmt.Printf("inf is a nil interface: %v\n", inf == nil)
	fmt.Printf("inf is a interface holding a nil pointer: %v\n", inf == (*Temp)(nil))
}
