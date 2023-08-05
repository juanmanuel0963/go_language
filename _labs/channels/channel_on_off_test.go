package channels

import (
	"fmt"
	"testing"
)

func TestOnOff(t *testing.T) {

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

	// Making a channel of value type string
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
