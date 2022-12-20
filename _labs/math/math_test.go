package math

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestMath_v1(t *testing.T) {

	num := 200

	sq_root_float := math.Sqrt(float64(num))
	fmt.Println(sq_root_float)

	sq_root_int := int(math.Sqrt(float64(num)))
	fmt.Println(sq_root_int)

	rand_int := rand.Intn(100)
	fmt.Println(rand_int)
}

func TestMath_v2(t *testing.T) {

	const Pi1 = 3.14
	fmt.Println(Pi1)

	const Pi2 = math.Pi
	fmt.Println(Pi2)
}
