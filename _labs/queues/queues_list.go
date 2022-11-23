package queues

import (
	"container/list"
	"fmt"
)

type QueueList struct {
	q *list.List
}

func (ql QueueList) Empty() bool {

	fmt.Println(ql.q)

	is_empty := false

	if ql.q == nil {
		is_empty = true
	} else if ql.q.Len() == 0 {
		is_empty = true
	}

	return is_empty
}

func (ql QueueList) Enqueue(v int) {

	ql.q.PushBack(v)
}

func (ql QueueList) Dequeue() any {

	myvar := ql.q.Remove(ql.q.Front())

	return myvar
}
