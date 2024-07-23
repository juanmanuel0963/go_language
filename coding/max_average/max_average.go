package main

import "fmt"

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func findMaxAverage(nums []int, k int) float64 {

	boxesLen := len(nums) - k + 1
	boxes := make([]int, len(nums))
	boxesAverage := make([]float64, len(nums))
	var maxAverage float64

	for i := 0; i < len(nums); i++ {

		limInf := i - k + 1
		limSup := i
		if limInf < 0 {
			limInf = 0
		}

		//fmt.Println(limInf)
		//fmt.Println(limSup)

		for j := limInf; j <= limSup; j++ {
			boxes[j] += nums[i]
			boxesAverage[j] = float64(boxes[j]) / float64(k)

			if boxesAverage[j] > maxAverage && j <= boxesLen {
				maxAverage = boxesAverage[j]
			}
		}

		//fmt.Println("-----")
	}

	//fmt.Println(boxes)
	//fmt.Println(boxesLen)
	//fmt.Println(boxesAverage)

	return maxAverage
}

func main() {

	//nums := []int{1, 12, -5, -6, 50, 3}
	//k := 4

	nums := []int{5}
	k := 1

	got := findMaxAverage(nums, k)
	fmt.Println(got)
}
