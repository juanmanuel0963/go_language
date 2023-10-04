package main

import "fmt"

//Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.

//Example

//For inputArray = [3, 6, -2, -5, 7, 3], the output should be
//solution(inputArray) = 21.
//7 and 3 produce the largest product.

func solution(inputArray []int) int {

	var operations []int

	for i := 0; i < len(inputArray)-1; i++ {
		calc := inputArray[i] * inputArray[i+1]
		operations = append(operations, calc)
	}

	fmt.Println(operations)

	max := 0

	for i, val := range operations {

		if i == 0 || val > max {
			max = val
		}
	}

	fmt.Println(max)

	return max
}

func main() {
	got := solution([]int{3, 6, -2, -5, 7, 3})
	fmt.Println(got)
}
