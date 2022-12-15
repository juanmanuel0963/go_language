package interfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// You can compare two interfaces using the == operator if the underlying types are “simple” and identical.
// Otherwise, the code will panic at runtime:
func TestCompare(t *testing.T) {

	var a interface{}
	var b interface{}

	a = 10
	b = 10

	assert.Equal(t, true, a == b)
}
