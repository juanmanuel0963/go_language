package stacks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStacksPointer(t *testing.T) {

	stack := StackPointer{}
	stack.PushPointer(10)
	stack.PushPointer(20)
	stack.PushPointer(30)
	stack.PushPointer(40)
	stack.PushPointer(50)
	assert.Equal(t, 5, len(stack))

	iPop := stack.PopPointer()
	assert.Equal(t, 50, iPop)

	fmt.Print("Poped: ")
	fmt.Println(iPop)

	assert.Equal(t, 4, len(stack))

	my_stack := stack.GetValuesPointer()
	fmt.Print("my_stack: ")
	fmt.Println(my_stack)

	assert.Equal(t, []int{10, 20, 30, 40}, stack.GetValuesPointer())

	assert.Equal(t, false, stack.EmptyPointer())

	i := 0
	for _, e := range stack {

		to_print := fmt.Sprintf("[%v]: %v", i, e)
		fmt.Println(to_print)
		i = i + 1
	}
}
