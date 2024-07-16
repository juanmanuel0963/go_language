package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {

	var maxLen int = 0
	var mergeStrings string

	if len(word1) >= len(word2) {
		maxLen = len(word1)
	} else {
		maxLen = len(word2)
	}

	for i := 0; i < maxLen; i++ {

		if i < (len(word1)) {
			mergeStrings = mergeStrings + string(word1[i])
		}

		if i < (len(word2)) {
			mergeStrings = mergeStrings + string(word2[i])
		}
	}
	return mergeStrings
}

func main() {
	got := mergeAlternately("a", "pqr")
	fmt.Println(got)
}
