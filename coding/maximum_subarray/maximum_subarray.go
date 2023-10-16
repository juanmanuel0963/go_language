package main

import "fmt"

func maxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	sum := 0
	var sumAll []int

	for i := 0; i < len(nums); i++ {

		sum = 0

		for j := 2; j <= len(nums); j++ {

			sum = 0

			for k := i; k < j; k++ {

				sum = sum + nums[k]

			}

			//fmt.Println(sum)
			sumAll = append(sumAll, sum)
		}

	}

	return findMax(sumAll)
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {

	got := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(got)

	got1 := maxSubArray([]int{5, 4, -1, 7, 8})
	fmt.Println(got1)
}
