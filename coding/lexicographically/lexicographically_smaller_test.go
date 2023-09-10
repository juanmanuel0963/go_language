package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution(t *testing.T) {

	count := lexicographically_smaller("ab12c", "1zz456")
	assert.Equal(t, 1, count)

	count = lexicographically_smaller("ab12c", "ab24z")
	assert.Equal(t, 3, count)

	count = lexicographically_smaller("96726", "9z34c")
	assert.Equal(t, 8, count)
}
