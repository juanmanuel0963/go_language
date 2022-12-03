package channels

import (
	"fmt"
	"testing"
)

func TestCalcChannels(t *testing.T) {

	number := 2

	// Making a channel of value type string
	sqrch := make(chan int)

	// Making a channel of value type string
	cubech := make(chan int)

	// Starting a concurrent goroutine
	go CalcSquares(number, sqrch)

	// Starting a concurrent goroutine
	go CalcCubes(number, cubech)

	// Receive a message from the channels
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

func TestChannels(t *testing.T) {

	////-------DirecTVChannel--------------

	// Making a channel of value type string
	myDirecTVChannel := make(chan string)

	// Starting a concurrent goroutine
	go DirecTVChannel(myDirecTVChannel)

	//-------Send the data to a channel--------------
	requestMessage := "TURN_ON"
	myDirecTVChannel <- requestMessage

	//-------Receive a message from the channel------
	responseMessage := <-myDirecTVChannel
	fmt.Println(responseMessage)

	////-------NetflixChannel--------------

	// Making a channel of value type struct
	myNetflixChannel := make(chan string)

	// Starting a concurrent goroutine
	go NetflixChannel(myNetflixChannel)

	//-------Send the data to a channel--------------
	requestMessageNetflixChannel := "TURN_ON"
	myNetflixChannel <- requestMessageNetflixChannel

	//-------Receive a message from the channel------
	responseMessageNetflixChannel := <-myNetflixChannel
	fmt.Println(responseMessageNetflixChannel)

}
