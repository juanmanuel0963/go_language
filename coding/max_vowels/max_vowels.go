package main

import "fmt"

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func maxVowels(s string, k int) int {

	var maxVowelCount = 0
	for i := 0; i <= len(s)-k; i++ {

		var vowelCount = 0

		for j := i; j < i+k; j++ {

			//fmt.Println(string(s[j]))

			if string(s[j]) == "a" || string(s[j]) == "e" || string(s[j]) == "i" || string(s[j]) == "o" || string(s[j]) == "u" ||
				string(s[j]) == "A" || string(s[j]) == "E" || string(s[j]) == "I" || string(s[j]) == "O" || string(s[j]) == "U" {
				vowelCount += 1
			}

			if vowelCount > maxVowelCount {
				maxVowelCount = vowelCount
			}
		}
		//fmt.Println("----")
	}

	return maxVowelCount
}

func main() {
	var s = "abciiidef"
	var k = 3

	got := maxVowels(s, k)
	fmt.Println(got)
}
