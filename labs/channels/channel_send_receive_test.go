package channels

import (
	"fmt"
	"testing"
)

func TestSendReceive(t *testing.T) {

	// declaring a bidirectional channel that transports data of type int
	c := make(chan int)
	fmt.Printf("%T\n", c)

	// starting the goroutine
	go SendReceive(10, c)

	//receiving data from the channel
	n := <-c
	fmt.Println("Value received:", n)

	fmt.Println("Exiting main ...")

}
