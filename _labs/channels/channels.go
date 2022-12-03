package channels

import "fmt"

func CalcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum

	//-------Closes the channel.
	close(squareop)
}

func CalcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum

	//-------Closes the channel.
	close(cubeop)
}

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
