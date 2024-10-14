/*
You are given a list of words. Write a function in GoLang called AnagramGroups
that groups together words that are anagrams of each other into separate lists.
Each group of anagrams should be returned as a slice of slices of strings.
If there are no anagrams in the list, the function should return an empty slice.
*/
package main

import (
	"fmt"
	"sort"
)

func Sort(text string) string {

	//eat
	//aet

	//tea
	//aet

	//ate
	//aet

	runes := []rune(text)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func AnagramGroups(words []string) [][]string {

	//Result Array of arrays
	result := make([][]string, 0)
	anagrams := make(map[string][]string)

	//Iterar array words
	for _, word := range words {
		sortedValue := Sort(word)
		//fmt.Println(sortedValue)
		anagrams[sortedValue] = append(anagrams[sortedValue], word)
	}

	fmt.Println(anagrams)

	for _, group := range anagrams {

		//fmt.Println(group)
		result = append(result, group)
	}

	return result

}

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("Anagram Groups:", AnagramGroups(words)) // Output: [["bat"] ["tan" "nat"] ["eat" "tea" "ate"]]
}
