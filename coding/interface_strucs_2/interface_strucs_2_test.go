package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution(t *testing.T) {

	response := solution([]string{"1*5", "1+2"})
	assert.Equal(t, 3, response[0])
	assert.Equal(t, 5, response[1])
}
