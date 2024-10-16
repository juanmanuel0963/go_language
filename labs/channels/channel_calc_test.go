package channels

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {

	number := 2

	// Making a channel of value type int
	sqrch := make(chan int)

	// Making a channel of value type int
	cubech := make(chan int)

	// Starting a concurrent goroutine
	go CalcSquares(number, sqrch)

	// Starting a concurrent goroutine
	go CalcCubes(number, cubech)

	// Receive a message from the channels
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("squares: ", squares)
	fmt.Println("cubes: ", cubes)
	fmt.Println("Final output: ", squares+cubes)
}
