package queues

import (
	"container/list"
)

type QueueList struct {
	q *list.List
}

func (ql QueueList) Empty() bool {

	return ql.q.Len() == 0
}

func (ql QueueList) Enqueue(v int) {

	ql.q.PushBack(v)
}

func (ql QueueList) ProcessNext() any {

	myvar := ql.q.Front()
	ql.q.Remove(ql.q.Front())

	return myvar
}
