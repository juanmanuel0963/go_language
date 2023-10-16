package main

import "fmt"

func solution(inputString string) bool {
	charCount := make(map[rune]int)

	// Count the frequency of each character in the string
	for _, char := range inputString {
		charCount[char]++
	}

	// Count the number of characters with odd frequencies
	oddCount := 0
	for _, count := range charCount {
		if count%2 != 0 {
			oddCount++
		}
	}

	//fmt.Println(oddCount)

	// A string can be rearranged to form a palindrome if it has at most one character with an odd frequency
	return oddCount <= 1
}

func main() {
	got := solution("aabb")
	fmt.Println(got)
}
