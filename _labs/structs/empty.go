package structs

import "fmt"

type Lamp struct{}

func (l Lamp) On() {
	println("On")

}

func (l Lamp) Off() {
	println("Off")
}

func WorkerRoutine(ch chan struct{}) {
	// Receive a message from the main program.
	<-ch
	fmt.Println("Signal Received")

	// Send a message to the main program.
	close(ch)
}
