package miscellaneous

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumericConversion(t *testing.T) {

	i := 55           //int
	j := 67.8         //float64
	sum := i + int(j) //j is converted to int
	fmt.Println(sum)

	assert.Equal(t, 122, sum)
}
