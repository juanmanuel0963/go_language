package stacks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStacks(t *testing.T) {

	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Push(50)
	assert.Equal(t, 5, len(s))

	iPop := s.Pop()
	assert.Equal(t, 50, iPop)
	fmt.Print(iPop)

	assert.Equal(t, 4, len(s))

	my_stack := s.GetValues()
	fmt.Println(my_stack)

	assert.Equal(t, []int{10, 20, 30, 40}, s.GetValues())

	assert.Equal(t, false, s.Empty())
}
