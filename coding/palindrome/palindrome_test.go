package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsPalindrome(t *testing.T) {

	isPalindrome1 := IsPalindrome("JUAN")
	assert.Equal(t, false, isPalindrome1)

	isPalindrome2 := IsPalindrome("TAT")
	assert.Equal(t, true, isPalindrome2)

}
