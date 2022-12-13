package channels

import (
	"fmt"
	"testing"
)

// Channel in both directions
func TestSum_v1(t *testing.T) {

	//Channel in both directions
	cInOut := make(chan int) //Receives one value at same time. It works
	//Channel running in background
	go Sum(cInOut)
	//Send value 1 down the channel
	cInOut <- 10
	//Send value 2 down the channel
	cInOut <- 15
	//Receives the addition
	r := <-cInOut
	//Prints the addition
	fmt.Println(r)

}

func TestSum_v2(t *testing.T) {

	//Channel in both directions
	cInOut := make(chan int, 2) //Expects 2 values in the channel. It doesn't works
	//Channel running in background
	go Sum(cInOut)
	//Send value 1 down the channel
	cInOut <- 10
	//Send value 2 down the channel
	cInOut <- 15
	//Receives the addition
	r := <-cInOut
	//Prints the addition
	fmt.Println(r)

}

// https://www.youtube.com/watch?v=rQQcIGqp0OU
func TestSumInOut_v3(t *testing.T) {

	//Channel input
	cIn := make(chan int, 8) //Expects 8 values in the channel
	//Channel output
	cOut := make(chan int, 4) //Outputs 4 values.
	//Channel running in background
	go SumMultipleInOut(cIn, cOut)
	//Send value 1 down the channel
	cIn <- 10
	//Send value 2 down the channel
	cIn <- 15
	//Send value 3 down the channel
	cIn <- 99
	//Send value 4 down the channel
	cIn <- 1
	//Send value 5 down the channel
	cIn <- 23
	//Send value 6 down the channel
	cIn <- 7
	//Send value 7 down the channel
	cIn <- 151
	//Send value 8 down the channel
	cIn <- 9

	//Receives the addition
	r := <-cOut
	//Prints the addition
	fmt.Println(r)

	//Receives the addition
	r = <-cOut
	//Prints the addition
	fmt.Println(r)

	//Receives the addition
	r = <-cOut
	//Prints the addition
	fmt.Println(r)

	//Receives the addition
	r = <-cOut
	//Prints the addition
	fmt.Println(r)
}

func TestIsPrime_v4(t *testing.T) {

	msg := PrimeMsg{42, false}

	//Create channels
	cIn := make(chan PrimeMsg, TEST_LEN)
	cOut := make(chan PrimeMsg, TEST_LEN)

	//Create workers (threads/Go Routines for concurrency)
	for i := 0; i < NUM_WORKERS; i++ {
		fmt.Println("Create workers: ", i)
		go IsPrime(cIn, cOut)
	}

	//Fill the input queue
	for i := 0; i < TEST_LEN; i++ {
		fmt.Println("Fill the input queue: ", i)
		//msg.num = rand.Intn(100) + 100
		msg.num = i
		cIn <- msg
	}

	//Read the answers
	for i := 0; i < TEST_LEN; i++ {
		//fmt.Println("Read the answers: ", i)
		msg = <-cOut
		fmt.Println(msg.num, "Is prime: ", msg.isPrime)
	}
}
