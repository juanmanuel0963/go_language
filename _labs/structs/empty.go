package structs

import "fmt"

type Lamp struct{}

func (l Lamp) On() {
	println("On")

}
func (l Lamp) Off() {
	println("Off")
}
func worker(ch chan struct{}) {
	// Receive a message from the main program.
	<-ch
	fmt.Println("roger2")

	// Send a message to the main program.
	close(ch)
}
