package main

import "fmt"

func solution(a, b []int) bool {
	// Check if the lengths of the two arrays are different
	if len(a) != len(b) {
		return false
	}

	// Create slices to store the indices of different elements
	var diffA, diffB []int

	// Iterate through the arrays and compare elements
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffA = append(diffA, a[i])
			diffB = append(diffB, b[i])
		}
	}

	if len(diffA) == 0 {
		return true
	}

	// If there are exactly two different elements in each array, and they can be swapped to match,
	// return true; otherwise, return false
	return len(diffA) == 2 && diffA[0] == diffB[1] && diffA[1] == diffB[0]
}

func main() {

	got1 := solution([]int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println(got1)
	/*
		got2 := solution([]int{1, 2, 3}, []int{2, 1, 3})
		fmt.Println(got2)

		got3 := solution([]int{1, 2, 2}, []int{2, 1, 1})
		fmt.Println(got3)
	*/
}
