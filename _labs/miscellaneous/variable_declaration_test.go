package miscellaneous

import (
	"fmt"
	"testing"
)

func TestVariableDeclaration(t *testing.T) {
	var a, b, c = 3, 4, "foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
