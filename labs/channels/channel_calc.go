package channels

import "fmt"

func CalcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		fmt.Println("squares number 1: ", number)

		digit := number % 10
		fmt.Println("squares digit: ", digit)

		sum += digit * digit
		number /= 10

		fmt.Println("squares number 2: ", number)
	}
	squareop <- sum

	//-------Closes the channel.
	close(squareop)
}

func CalcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		fmt.Println("cubes number 1: ", number)
		digit := number % 10

		fmt.Println("cubes digit: ", digit)

		sum += digit * digit * digit
		number /= 10

		fmt.Println("cubes number 2: ", number)
	}
	cubeop <- sum

	//-------Closes the channel.
	close(cubeop)
}
