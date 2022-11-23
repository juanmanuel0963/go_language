package queues

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueues_v1(t *testing.T) {

	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)

	assert.Equal(t, 5, len(q))

	v := q.Dequeue()
	fmt.Println(v)
	assert.Equal(t, 4, len(q))

	assert.Equal(t, false, q.Empty())
}
