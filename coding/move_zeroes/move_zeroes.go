package main

import "fmt"

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func moveZeroes(nums []int) []int {

	counter := 0
	response := make([]int, 0)

	for i := 0; i < len(nums); i++ {

		if nums[i] == 0 {
			counter += 1
		} else {
			response = append(response, nums[i])
		}
	}

	for i := 0; i < counter; i++ {
		response = append(response, 0)
	}

	return response
}

func main() {

	nums := []int{0, 1, 0, 3, 12}

	got := moveZeroes(nums)
	fmt.Println(got)
}
