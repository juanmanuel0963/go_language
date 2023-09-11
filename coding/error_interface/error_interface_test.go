// CODE EXAMPLE VALID FOR COMPILING

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getError(t *testing.T) {

	var ce1 MyError
	ce1.Msg = "Error #1"
	val1 := getError(ce1)
	assert.Equal(t, "Error #1", val1)

	ce2 := MyError{Msg: "Error #2"}
	val2 := getError(ce2)
	assert.Equal(t, "Error #2", val2)
}
