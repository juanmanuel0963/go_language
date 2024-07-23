package main

import (
	"fmt"
)

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func containsValue(nums []int, k int) bool {

	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			return true
		}
	}

	return false
}

func setFoundValue(aControl []int, k int) []int {

	for i := 0; i < len(aControl); i++ {
		if aControl[i] == k {
			//fmt.Println(aControl[i])
			aControl[i] = -1
			break
		}
	}

	return aControl
}

func maxOperations(nums []int, k int) int {

	//fmt.Println(k)
	aControl := make([]int, 0)
	pairCounter := 0
	for i := 0; i < len(nums); i++ {
		pair := k - nums[i]
		//fmt.Println("nums[i]: " + strconv.Itoa(nums[i]) + "-------")
		//fmt.Println("pair: " + strconv.Itoa(pair))
		if pair > 0 {
			//Check if someone is searching me
			searchingMe := containsValue(aControl, nums[i])
			if searchingMe {
				//fmt.Println("New pair found 1: " + strconv.Itoa(nums[i]))
				aControl = setFoundValue(aControl, nums[i])
				pairCounter += 1
				//fmt.Println("-----")
			} else {
				//Add the pair for future search
				aControl = append(aControl, pair)
			}
		}
	}
	//fmt.Println(aControl)
	return pairCounter
}

func main() {

	//nums := []int{1, 2, 3, 4}
	//k := 5

	//nums := []int{3, 1, 3, 4, 3}
	//k := 6

	//nums := []int{3, 3, 3, 3, 3, 3}
	//k := 6

	//nums := []int{2, 4, 5, 3, 2}
	//k := 3

	nums := []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}
	k := 3

	got := maxOperations(nums, k)

	fmt.Println(got)

}
