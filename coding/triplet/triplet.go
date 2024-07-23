package main

import "fmt"

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func increasingTriplet(nums []int) bool {

	counter := 0

	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			counter += 1
		}
	}

	//fmt.Println(counter)
	//fmt.Println(nums)

	return counter > 0
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	got := increasingTriplet(nums)
	fmt.Println(got)
}
