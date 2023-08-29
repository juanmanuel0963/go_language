package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// Receives an array of words and returns a map with matching anagrams words
func findAnagrams(wordsDictionary []string) map[string][]string {
	//Declares matrix of chars. Each rows is a word of the dictionary
	var runeWords [][]rune
	//Declares map (hash) to save the word (key) and array of anagrams (value)
	anagrams := make(map[string][]string)
	//Step 1. We need to convert the array of words into a matrix of chars
	//Iterates each word in the dictionary
	for _, word := range wordsDictionary {
		//Array of chars for the current word
		var runeLetters []rune
		//Iterates the word and fetch each letter
		for _, letter := range word {
			//Adds the letter to the array of chars
			runeLetters = append(runeLetters, letter)
		}
		//Adds the word (as array of chars) to the matrix of words
		runeWords = append(runeWords, runeLetters)
		//Adds a new key to the anagrams map (hash)
		anagrams[string(word)] = nil
	}
	//Step 2. Once we have the words dictionary converted into a matrix of chars, we can take each word (row) and find the permutations
	//Iterates each row of the matrix
	for _, wordLetters := range runeWords {
		//Finds the permutations for the word. It returns a matrix and each row is a permutation
		permutations := permuteLetters(wordLetters)
		//For each permutation
		for _, perm := range permutations {
			//Checks if the permutations is equal to any word in the dictionary. It compares the permutation agains the map(hash) of dictionary words
			if myAnagrams, found := anagrams[string(perm)]; found {
				//if the permutations is not part of its parent word
				if string(wordLetters) != string(perm) {
					//If the permutation's parent word is not added yet to the anagrams list
					contains := slices.Contains(myAnagrams, string(wordLetters)) // true
					if !contains {
						//Adds the permutation's parent word to the array of anagrams for the found key (dictionary words)
						myAnagrams = append(myAnagrams, string(wordLetters))
						anagrams[string(perm)] = myAnagrams
					}
				}
			}
		}
	}

	//Returns a map (hash) with each dictionary word and the matching anagram words
	return anagrams
}

// Receives a word as an array of chars and returns a matrix with rows of permutations
func permuteLetters(letters []rune) [][]rune {
	var result [][]rune
	backtrack(letters, 0, &result)
	return result
}

// Finds the permutations. Only this function is courtesy of ChatGPT ;)
func backtrack(letters []rune, start int, result *[][]rune) {
	if start == len(letters) {
		current := make([]rune, len(letters))
		copy(current, letters)
		*result = append(*result, current)
		return
	}

	for i := start; i < len(letters); i++ {
		letters[start], letters[i] = letters[i], letters[start]
		backtrack(letters, start+1, result)
		letters[start], letters[i] = letters[i], letters[start]
	}
}

func main() {

	//Load the words dictionary into an array
	wordsDictionary := []string{"CAT", "ACT", "MEAT", "TEAM", "NIGHT", "THING", "THEEYES", "THEYSEE", "DOG"}

	//Send the dictionary to the function which returns a map (key:value / hash) with matching anagrams words
	var anagrams = findAnagrams(wordsDictionary)

	//Prints the words and related anagrams
	for key, value := range anagrams {
		fmt.Println(key+" : ", value)
	}
}
