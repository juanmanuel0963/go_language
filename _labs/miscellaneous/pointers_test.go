package miscellaneous

import (
	"fmt"
	"testing"
)

func TestPointers(t *testing.T) {

	x := 5
	p := &x
	q := *p

	fmt.Printf("p = %d", *p)
	fmt.Println()
	fmt.Printf("q = %d", q)
	fmt.Println()
}
