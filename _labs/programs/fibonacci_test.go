package programs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {

	n := 2

	fibonacciNumber := Fibonacci(n)

	fmt.Println(fibonacciNumber)

	assert.Equal(t, 0, Fibonacci(0))
	assert.Equal(t, 1, Fibonacci(1))
	assert.Equal(t, 1, Fibonacci(2))
	assert.Equal(t, 2, Fibonacci(3))
	assert.Equal(t, 3, Fibonacci(4))
	assert.Equal(t, 5, Fibonacci(5))
	assert.Equal(t, 8, Fibonacci(6))
	assert.Equal(t, 13, Fibonacci(7))
}
