package miscellaneous

import (
	"fmt"
	"math"
	"testing"
)

func TestMath(t *testing.T) {

	num := 200

	sq_root_float := math.Sqrt(float64(num))
	fmt.Println(sq_root_float)

	sq_root_int := int(math.Sqrt(float64(num)))
	fmt.Println(sq_root_int)
}
