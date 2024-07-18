package main

import (
	"fmt"
)

func reverseWords(s string) string {

	var sWord string
	var aWords []string

	//Iterate the input string and extract individual words into an slice
	for i := 0; i < len(s); i++ {

		if string(s[i]) != " " {
			sWord = sWord + string(s[i])
		}

		if string(s[i]) == " " && len(sWord) > 0 {
			aWords = append(aWords, sWord)
			sWord = ""
		}
	}

	if len(sWord) > 0 {
		aWords = append(aWords, sWord)
		sWord = ""
	}

	var sReverseWords string

	for i := len(aWords) - 1; i >= 0; i-- {
		sReverseWords = sReverseWords + aWords[i]

		if i > 0 {
			sReverseWords = sReverseWords + " "
		}
	}

	//fmt.Println(sReverseWords)

	return sReverseWords
}

func main() {
	got := reverseWords("  the   sky   is   blue  ")
	fmt.Println(got)
}

/*
func reverseWords(s string) string {

	//split the string with whitespace
	splitted := strings.Split(s, " ")
	//fmt.Println(splitted)

	var aWords []string
	var sWords string

	//add single words into a slice
	for i := len(splitted) - 1; i >= 0; i-- {

		if splitted[i] != "" {
			aWords = append(aWords, splitted[i])
		}
	}

	//concat words and single space
	for i := 0; i < len(aWords); i++ {
		sWords = sWords + aWords[i]

		if i < len(aWords)-1 {
			sWords = sWords + " "
		}
	}

	return sWords
}
*/
