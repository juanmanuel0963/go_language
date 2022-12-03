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

// When building an object, and only being interested in a grouping of methods and no intermediary data,
// or when you do not plan to retain the object state.
// In the example below it does not make a difference whether the method is called on the same (case #1)
// or on two different objects (case #2):
func TestEmptyStructs_v3(t *testing.T) {
	// Case #1.
	var lamp Lamp
	lamp.On()
	lamp.Off()
	// Output:
	// on
	// off
}

func TestEmptyStructs_v4(t *testing.T) {
	// Case #2.
	Lamp{}.On()
	Lamp{}.Off()
	// Output:
	// on
	// off
}

// When you need a channel to signal an event, but do not really need to send any data.
// This event is also not the last one in the sequence, in which case you would use the close(ch) built-in function.
func TestEmptyStructs_v5(t *testing.T) {
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
