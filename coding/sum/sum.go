package main

import "fmt"

//Write a function that returns the sum of two numbers.

//Example

//For param1 = 1 and param2 = 2, the output should be
//solution(param1, param2) = 3.

func solution(param1 int, param2 int) int {

	return param1 + param2
}

func main() {

	sum := solution(100, 50)
	fmt.Println(sum)
}
