package if_else

import (
	"fmt"
	"testing"
)

func TestIfElse(t *testing.T) {

	var x, y, result int

	x = 5
	y = 10

	if x > y {
		result = x
	} else {
		result = y
	}

	fmt.Println(result)

}
