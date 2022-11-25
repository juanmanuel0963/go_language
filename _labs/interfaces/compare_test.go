package interfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {

	var a interface{}
	var b interface{}

	a = 10
	b = 10

	assert.Equal(t, true, a == b)
}
