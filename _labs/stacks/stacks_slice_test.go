package stacks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStacks(t *testing.T) {

	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)
	assert.Equal(t, 5, len(stack))

	iPop := stack.Pop()
	assert.Equal(t, 50, iPop)
	fmt.Print(iPop)

	assert.Equal(t, 4, len(stack))

	my_stack := stack.GetValues()
	fmt.Println(my_stack)

	assert.Equal(t, []int{10, 20, 30, 40}, stack.GetValues())

	assert.Equal(t, false, stack.Empty())

	for _, e := range stack {
		fmt.Println(e)
	}
}
