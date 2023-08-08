package movingaverages

import (
	"github.com/juanmanuel0963/go_language/v1/advance/datastructures/queues"
)

// New creates a new moving average structure with a fixed sliding window
func New(windowSize int) *MovingAverage {
	return &MovingAverage{sum: 0, windowSize: windowSize, queue: queues.New[int]()}
}

// MovingAverage keeps track of the moving average
type MovingAverage struct {
	sum        int
	windowSize int
	queue      *queues.Queue[int]
}

// Add adds a number to the moving average calculations
func (ma *MovingAverage) Add(number int) {
	ma.sum += number
	ma.queue.Enqueue(number)
	if ma.queue.Size() <= ma.windowSize {
		return
	}

	value, _ := ma.queue.Dequeue()
	ma.sum -= value
}

// GetAverage calculates the moving average
func (ma *MovingAverage) GetAverage() float64 {
	return float64(ma.sum) / float64(ma.queue.Size())
}

// GetSum returns the sum within the sliding window
func (ma *MovingAverage) GetSum() int {
	return ma.sum
}
