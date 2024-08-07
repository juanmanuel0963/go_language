package main

import (
	"fmt"
	"strconv"
)

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func longestSubarray(nums []int) int {
	left := 0
	right := 0
	zeroCount := 0
	maxLength := 0
	//                0, 1, 2, 3, 4, 5, 6, 7, 8
	//	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	for right < len(nums) {
		//If current value is 0, increase the counter
		if nums[right] == 0 {
			zeroCount++
		}
		// If there are more than one zero, shrink the window from the left
		for zeroCount > 1 {
			fmt.Println("right: " + strconv.Itoa(right) + ", left: " + strconv.Itoa(left) + ", zeroCount: " + strconv.Itoa(zeroCount))
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// Calculate the length of the current window minus one element to be deleted
		maxLength = max(maxLength, right-left)

		fmt.Println("right: " + strconv.Itoa(right) + ", left: " + strconv.Itoa(left) + ", zeroCount: " + strconv.Itoa(zeroCount) + ", maxLength: " + strconv.Itoa(maxLength))

		right++
	}
	fmt.Println("Final zeroCount: " + strconv.Itoa(zeroCount))
	// Special case: If the entire array is 1s, we need to delete one element
	if maxLength == len(nums) {
		return maxLength - 1
	}
	return maxLength
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//nums := []int{1, 1, 0, 1}
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	//nums := []int{1, 1, 1}
	got := longestSubarray(nums)
	fmt.Println("got: " + strconv.Itoa(got)) // Output: 5
}

/*
func longestSubarray(nums []int) int {
	var maxConsecutiveCount = 0
	var consecutiveCount = 0

	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums); j++ {

			if i != j {
				if nums[j] == 1 {
					consecutiveCount += 1
					if consecutiveCount > maxConsecutiveCount {
						maxConsecutiveCount = consecutiveCount
					}
				} else if nums[j] == 0 {
					consecutiveCount = 0
				}
			}
		}
		consecutiveCount = 0
	}

	return maxConsecutiveCount
}

func main() {

	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}

	got := longestSubarray(nums)
	fmt.Println(got)
}
*/
