package main

import "fmt"

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func findDifference(nums1 []int, nums2 []int) [][]int {

	fmt.Println(nums1)
	fmt.Println(nums2)
	response := make([][]int, 0)
	temp1 := make([]int, 0)
	temp2 := make([]int, 0)
	countNotExist := 0

	for i := 0; i < len(nums1); i++ {
		countNotExist = 0
		for j := 0; j < len(nums2); j++ {
			if nums1[i] != nums2[j] {
				countNotExist += 1
			}
		}

		if countNotExist == len(nums2) {
			if !checkExist(temp1, nums1[i]) {
				temp1 = append(temp1, nums1[i])
			}
		}
	}

	response = append(response, temp1)

	for i := 0; i < len(nums2); i++ {
		countNotExist = 0
		for j := 0; j < len(nums1); j++ {
			if nums2[i] != nums1[j] {
				countNotExist += 1
			}
		}

		if countNotExist == len(nums1) {
			if !checkExist(temp2, nums2[i]) {
				temp2 = append(temp2, nums2[i])
			}
		}
	}

	response = append(response, temp2)

	return response
}

func checkExist(nums []int, toCheck int) bool {

	for i := 0; i < len(nums); i++ {
		if nums[i] == toCheck {
			//
			return true
		}
	}

	return false
}

func main() {
	//nums1 := []int{1, 2, 3}
	//nums2 := []int{2, 4, 6}
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}
	got := findDifference(nums1, nums2)
	fmt.Println(got)
}
