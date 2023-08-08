package stacks

import (
	"container/list"
)

type StackList struct {
	s *list.List
}

func (sl StackList) Empty() bool {

	return sl.s.Len() == 0
}

func (sl StackList) ToStack(v int) {

	sl.s.PushFront(v)
}

func (sl StackList) ProcessNext() any {

	myvar := sl.s.Front()
	sl.s.Remove(sl.s.Front())

	return myvar
}
