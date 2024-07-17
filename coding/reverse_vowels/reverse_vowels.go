package main

import (
	"fmt"
	"strings"
)

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func reverseVowels(s string) string {

	var vowels []string
	var word []string

	for i := 0; i < len(s); i++ {

		word = append(word, string(s[i]))

		if string(s[i]) == "a" || string(s[i]) == "e" || string(s[i]) == "i" || string(s[i]) == "o" || string(s[i]) == "u" || string(s[i]) == "A" || string(s[i]) == "E" || string(s[i]) == "I" || string(s[i]) == "O" || string(s[i]) == "U" {
			//fmt.Println(string(s[i]))
			vowels = append(vowels, string(s[i]))
		}
	}

	//fmt.Println(word)
	j := 0
	//fmt.Println(j)

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "a" || string(s[i]) == "e" || string(s[i]) == "i" || string(s[i]) == "o" || string(s[i]) == "u" || string(s[i]) == "A" || string(s[i]) == "E" || string(s[i]) == "I" || string(s[i]) == "O" || string(s[i]) == "U" {
			word[i] = vowels[j]
			j = j + 1
		}
	}

	//fmt.Println(vowels)
	//fmt.Println(word)

	str := strings.Join(word, "")

	return str
}

func main() {
	got := reverseVowels("hellO")
	fmt.Println(got)
}
