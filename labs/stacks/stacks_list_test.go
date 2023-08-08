package stacks

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStacksList_v1(t *testing.T) {

	stack := list.New()

	stack.PushFront(100)
	stack.PushFront(200)
	stack.PushFront(300)

	fmt.Println(stack.Len())
	assert.Equal(t, 3, stack.Len())

	//Process next element
	stack.Remove(stack.Front())
	assert.Equal(t, 2, stack.Len())

	for e := stack.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		fmt.Println(e.Value)
	}

}

func TestStacksList_v2(t *testing.T) {

	stack := StackList{}
	stack.s = list.New()

	stack.s.PushFront(100)
	stack.s.PushFront(200)
	stack.s.PushFront(300)

	fmt.Println(stack.s.Len())
	assert.Equal(t, 3, stack.s.Len())

	//Process next element
	stack.ProcessNext()

	fmt.Println(stack.s.Len())
	assert.Equal(t, 2, stack.s.Len())

	for e := stack.s.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		fmt.Println(e.Value)
	}
}
