package main

import (
	"fmt"
)

//Given matrix, a rectangular matrix of integers, where each value represents the cost of the room, your task is to return the total sum of all rooms that are suitable for the CodeBots (ie: add up all the values that don't appear below a 0).
//For

//matrix = [[0, 1, 1, 2],
//          [0, 5, 0, 0],
//          [2, 0, 3, 3]]
//the output should be
//solution(matrix) = 9.

func solution(matrix [][]int) int {

	fmt.Println(len(matrix[0]))
	fmt.Println(len(matrix))

	totalSum := 0
	//Going throught the matrix horizontally
	for i := 0; i < len(matrix[0]); i++ {
		sum := 0

		//Going throught the matrix vertically
		for j := 0; j < len(matrix); j++ {
			sum += matrix[j][i]

			if matrix[j][i] > 0 {
				totalSum += matrix[j][i]
			} else {
				break
			}
		}

		fmt.Printf("Sum of elements in column %d is %d\n", i+1, sum)
	}

	fmt.Printf("Total Sum of elements is %d\n", totalSum)

	//If value is > 0 then Sum the position
	//If not then break

	return totalSum
}

func main() {
	got := solution([][]int{{0, 1, 1, 2}, {0, 5, 0, 0}, {2, 0, 3, 3}})
	fmt.Println(got)
}
