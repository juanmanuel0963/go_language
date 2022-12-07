package structs

import "fmt"

func WorkerRoutine(ch chan struct{}) {
	// Receive a message from the main program.
	<-ch
	fmt.Println("Signal Received")

	// Send a message to the main program.
	close(ch)
}
