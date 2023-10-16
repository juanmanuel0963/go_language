package main

import "fmt"

func solution(inputArray []int) int {

	moves := 0

	for i := 1; i < len(inputArray); i++ {

		if inputArray[i] <= inputArray[i-1] {

			diff := inputArray[i-1] - inputArray[i]

			inputArray[i] = inputArray[i] + diff + 1

			moves = moves + diff + 1
		}
	}

	fmt.Println(inputArray)

	return moves
}

func main() {
	got := solution([]int{1, 1, 1})
	fmt.Println(got)
}
