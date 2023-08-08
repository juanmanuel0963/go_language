package stacks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStacks(t *testing.T) {

	stack := Stack{}
	stack = stack.Push(10)
	stack = stack.Push(20)
	stack = stack.Push(30)
	stack = stack.Push(40)
	stack = stack.Push(50)
	assert.Equal(t, 5, len(stack))

	stack, iPop := stack.Pop()
	assert.Equal(t, 50, iPop)

	fmt.Print("Poped: ")
	fmt.Println(iPop)

	assert.Equal(t, 4, len(stack))

	my_stack := stack.GetValues()
	fmt.Print("my_stack: ")
	fmt.Println(my_stack)

	assert.Equal(t, []int{10, 20, 30, 40}, stack.GetValues())
	assert.Equal(t, false, stack.Empty())

	i := 0
	for _, e := range stack {

		to_print := fmt.Sprintf("[%v]: %v", i, e)
		fmt.Println(to_print)
		i = i + 1
	}
}
