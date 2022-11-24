package queues

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueuesList_v1(t *testing.T) {

	queue := list.New()

	queue.PushBack(100)
	queue.PushBack(200)
	queue.PushBack(300)

	assert.Equal(t, 3, queue.Len())
	fmt.Println(queue.Len())

	//Process next element
	queue.Remove(queue.Front())
	assert.Equal(t, 2, queue.Len())

	for e := queue.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		fmt.Println(e.Value)
	}

}

func TestQueuesList_v2(t *testing.T) {

	queue := QueueList{}
	queue.q = list.New()

	fmt.Println(queue.Empty())

	queue.Enqueue(100)
	queue.Enqueue(200)
	queue.Enqueue(300)

	assert.Equal(t, 3, queue.q.Len())
	fmt.Println(queue.Empty())

	queue.ProcessNext()
	assert.Equal(t, 2, queue.q.Len())
	fmt.Println(queue.q.Len())

	for e := queue.q.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		fmt.Println(e.Value)
	}
}
