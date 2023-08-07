package queues

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueues_v1(t *testing.T) {

	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Enqueue(50)

	assert.Equal(t, 5, len(queue))

	v := queue.Dequeue()
	fmt.Println(v)
	fmt.Println(" ")

	assert.Equal(t, 4, len(queue))

	assert.Equal(t, false, queue.Empty())

	for _, e := range queue {
		fmt.Println(e)
	}

}
