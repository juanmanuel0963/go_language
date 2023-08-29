package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"golang.org/x/exp/slices"
)

func Test_findAnagrams(t *testing.T) {

	//Load the words dictionary into an array
	wordsDictionary := []string{"CAT", "ACT", "MEAT", "TEAM", "NIGHT", "THING", "THEEYES", "THEYSEE", "DOG"}

	//Send the dictionary to the function which returns a map (key:value / hash) with matching anagrams words
	var anagrams = findAnagrams(wordsDictionary)

	if myAnagrams, found := anagrams["CAT"]; found {

		contains := slices.Contains(myAnagrams, "ACT")

		assert.Equal(t, contains, true)
	}

}
