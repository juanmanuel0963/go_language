package miscellaneous

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeConversion(t *testing.T) {
	//Float To Int
	i := 55                //int
	j := 67.8              //float64
	float_to_int := int(j) //j is converted to int
	fmt.Println(float_to_int)
	sum := i + float_to_int
	assert.Equal(t, 122, sum)

	//Int to String
	int_to_str := strconv.Itoa(55) //i is converted to string
	fmt.Println(int_to_str)
	assert.Equal(t, "55", int_to_str)

}
