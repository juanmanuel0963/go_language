package channels

import (
	"fmt"
)

func DirecTVChannel(myDirecTVChannel chan string) {

	// Receive a message from the main program.
	requestMessage := <-myDirecTVChannel
	fmt.Println(requestMessage)

	//-------Send the data to a channel--------------
	responseMessage := "DIRECTV_TURNED_ON"
	myDirecTVChannel <- responseMessage

	//-------Closes the channel.
	close(myDirecTVChannel)
}

func NetflixChannel(myNetflixChannel chan string) {

	// Receive a message from the main program.
	requestMessage := <-myNetflixChannel
	fmt.Println(requestMessage)

	//-------Send the data to a channel--------------
	responseMessage := "NETFLIX_TURNED_ON"
	myNetflixChannel <- responseMessage

	//-------Closes the channel.
	close(myNetflixChannel)
}
