package channels

import (
	"fmt"
	"math"
)

// Channel in both directions
func Sum(cInOut chan int) {
	//blocks the channel until something comes
	x := <-cInOut
	//blocks the channel until something comes
	y := <-cInOut
	//returns the addition
	cInOut <- x + y
}

// https://www.youtube.com/watch?v=rQQcIGqp0OU
// Channel in both directions
func SumMultipleInOut(cIn chan int, cOut chan int) {
	for {
		//blocks the channel until something comes
		x := <-cIn
		//blocks the channel until something comes
		y := <-cIn

		//prints the addition
		fmt.Println(x, "+", y, "=", x+y)

		//returns the addition
		cOut <- x + y
	}
}

type PrimeMsg struct {
	num     int
	isPrime bool
}

const TEST_LEN = 100
const NUM_WORKERS = 5

func IsPrime(cIn chan PrimeMsg, cOut chan PrimeMsg) {

	//id := rand.Intn(100)
	i := 0
	for {
		msg := <-cIn
		num := msg.num
		//fmt.Println(id, "is testing", num)
		sq_root := int(math.Sqrt(float64(num)))
		for i = 2; i <= sq_root; i++ {
			if num%i == 0 {
				msg.isPrime = false
				cOut <- msg
				break
			}
		}
		if i > sq_root {
			msg.isPrime = true
			cOut <- msg
		}
	}
}
