package main

import "fmt"

func productExceptSelf(nums []int) []int {

	n := len(nums)
	answer := make([]int, n)

	// Initialize the answer array where answer[i] contains the product of all the elements to the left of nums[i]
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	fmt.Println(answer)

	// R will contain the product of all the elements to the right of nums[i]
	R := 1
	for i := n - 1; i >= 0; i-- {
		// answer[i] now contains the product of all the elements to the left and right of nums[i]
		answer[i] = answer[i] * R
		R *= nums[i]
	}
	return answer
}

//nums	 1	2	3	4
//Left	 1	1	2	6
//Right	 24	12	4	1
//Answer 24	12	8	6

func main() {

	//nums := []int{1, 2, 3, 4}
	nums := []int{-1, 1, 0, -3, 3}
	got := productExceptSelf(nums)

	fmt.Println(got)
}
