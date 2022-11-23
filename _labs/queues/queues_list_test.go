package queues

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueuesList(t *testing.T) {

	ql := QueueList{}
	q := list.New()
	ql.q = q

	fmt.Println(ql.q)

	assert.Equal(t, true, ql.Empty())

	ql.Enqueue(10)

	assert.Equal(t, false, ql.Empty())

	fmt.Println(ql.q)
	/*	ql.Enqueue(20)
		ql.Enqueue(30)
		ql.Enqueue(40)
		ql.Enqueue(50)
	*/
	//assert.Equal(t, 5, ql.q.Len())
}
