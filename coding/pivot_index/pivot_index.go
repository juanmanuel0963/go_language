package main

import (
	"fmt"
)

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func pivotIndex(nums []int) int {
	fmt.Println(nums)
	leftSum := 0
	rightSum := 0
	pivotIndex := -1
	totalRightSum := 0

	for i := 1; i < len(nums); i++ {
		totalRightSum += nums[i]
	}

	for i := 0; i < len(nums); i++ {

		if i == 0 {
			leftSum = 0
			rightSum = totalRightSum
		} else if i == len(nums) {
			leftSum += nums[i-1]
			rightSum = 0
		} else {
			leftSum += nums[i-1]
			totalRightSum = totalRightSum - nums[i]
			rightSum = totalRightSum
		}

		//fmt.Println("Value: " + strconv.Itoa(nums[i]) + ", leftSum: " + strconv.Itoa(leftSum) + ", rightSum: " + strconv.Itoa(rightSum))

		if leftSum == rightSum {
			pivotIndex = i
			break
		}

		rightSum = 0
	}

	return pivotIndex
}

func main() {

	nums := []int{1, 7, 3, 6, 5, 6}
	//nums := []int{2, 1, -1}
	//nums := []int{1, 2, 3}
	got := pivotIndex(nums)
	fmt.Println(got)
}
