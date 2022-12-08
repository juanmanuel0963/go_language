package strings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringContains(t *testing.T) {

	contains := StringContains("Welcome to New York City", "City")
	fmt.Println(contains)

	assert.Equal(t, true, StringContains("Welcome to New York City", "City"))
	assert.Equal(t, true, StringContains("Displaying the result", "result"))

}
