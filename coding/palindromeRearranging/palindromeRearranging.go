package main

import "fmt"

func solution(inputString string) bool {

	swap := make([]string, len(inputString))
	var copyInputString []string

	for i := 0; i < len(inputString); i++ {
		copyInputString = append(copyInputString, string(inputString[i]))
	}

	for i := 0; i < len(inputString); i++ {

		for j := 0; j < len(inputString); j++ {

			copy(swap, copyInputString)

			//Creating permutation
			if i != j {
				swap[i] = string(inputString[j])
				swap[j] = string(inputString[i])
			}

			if isPalindrome(swap) {
				return true
			}

			//fmt.Printf("permutation: %s IsPalindrome: %v \n", swap, bIsPalindrome)

		}

		//fmt.Println("-----")
	}

	return false
}

func isPalindrome(inputString []string) bool {

	for i, j := 0, len(inputString)-1; i < j; i, j = i+1, j-1 {

		if inputString[i] != inputString[j] {
			return false
		}

	}

	return true
}

func main() {
	got := solution("aabb")
	fmt.Println(got)
}
