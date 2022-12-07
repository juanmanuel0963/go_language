package structs

import (
	"fmt"
	"testing"
	"unsafe"
)

// You would use an empty struct when you would want to save some memory.
// Empty structs do not take any memory for its value.
func TestEmptyStructs_v1(t *testing.T) {

	a := struct{}{}
	println(unsafe.Sizeof(a))
	// Output: 0
}

// When implementing a data set:
func TestEmptyStructs_v2(t *testing.T) {
	set := make(map[string]struct{})
	for _, value := range []string{"apple", "orange", "strawberry"} {
		set[value] = struct{}{}
	}
	fmt.Println(set)
	// Output: map[orange:{} apple:{}]
}

// When you need a channel to signal an event, but do not really need to send any data.
// This event is also not the last one in the sequence, in which case you would use the close(ch) built-in function.
func TestEmptyStructs_v6(t *testing.T) {
	// Making a channel of value type struct
	ch := make(chan struct{})

	// Starting a concurrent goroutine
	go WorkerRoutine(ch)

	// Send signal to worker goroutine
	ch <- struct{}{}

	// Receive a message from the workerRoutine.
	<-ch
	fmt.Println("Signal Received")
}
